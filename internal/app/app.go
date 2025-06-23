package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"taskapi/internal/logger"
	"taskapi/internal/scr/tasks"
	"taskapi/internal/services"
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

	//port := envs["LBAPI_LISTEN_PORT"]
	srv := &http.Server{
		Addr:         ":" + "8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      router,
	}

	return srv, nil
}

func Run(ctx context.Context) error {
	log.Infof("App is running... on port %s", ":8080")

	srv, err := newServer()
	if err != nil {
		return err
	}

	runPeriodicJobs(ctx)
	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = srv.Shutdown(shutdownCtx)
		if err != nil {
			log.Errorf("Graceful shutdown failed: %s", err.Error())
		}

		log.Infof("App stopped gracefully")
	}()

	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("can't start http server: %s", err)
	}

	return nil
}
