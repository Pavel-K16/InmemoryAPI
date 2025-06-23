package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"taskapi/internal/logger"
	"taskapi/internal/scr/tasks"
	"taskapi/internal/services"
	"taskapi/internal/utils"
	"time"

	"github.com/gorilla/mux"
)

var (
	log = logger.LoggerInit()
)

func runPeriodicJobs(ctx context.Context) {
	go services.CacheTasksWatcher(ctx) // watcher for tasks
}

func newRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks/{id}", tasks.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", tasks.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", tasks.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks", tasks.GetAllTasks).Methods("GET")

	return router
}

func newServer() (*http.Server, error) {
	router := newRouter()

	envs := utils.GetEnvs()

	if err := utils.SetEnvs(envs); err != nil {
		return nil, err
	}

	port := envs["TASKAPI_LISTEN_PORT"]
	readTimeout := envs["TASKAPI_HTTP_READ_TIMEOUT_SEC"]
	writeTimeout := envs["TASKAPI_HTTP_WRITE_TIMEOUT_SEC"]

	readTimeoutSec, err := strconv.Atoi(readTimeout)
	if err != nil {
		return nil, err
	}

	writeTimeoutSec, err := strconv.Atoi(writeTimeout)
	if err != nil {
		return nil, err
	}

	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  time.Duration(readTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(writeTimeoutSec) * time.Second,
		Handler:      router,
	}

	return srv, nil
}

func Run(ctx context.Context) error {
	srv, err := newServer()
	if err != nil {
		return err
	}

	port := os.Getenv("TASKAPI_LISTEN_PORT")

	log.Infof("App is running... on port :%s", port)

	runPeriodicJobs(ctx)

	go func() error {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("can't start http server: %s", err)
		}

		return nil
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = srv.Shutdown(shutdownCtx)
	if err != nil {
		log.Errorf("Graceful shutdown failed: %s", err.Error())
	}

	log.Infof("App stopped gracefully")

	return nil
}
