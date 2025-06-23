package services

import (
	"context"
	"os"
	"strconv"
	"sync"
	e "taskapi/internal/entities"
	"taskapi/internal/logger"
	"taskapi/internal/tools"
	"time"
)

var (
	log = logger.LoggerInit()
)

func CacheTasksWatcher(ctx context.Context) {
	log.Infof("Service CacheTasksWatcher is running...")

	interval := os.Getenv("SYNC_TASKSTATUS_INTERVAL_SEC")
	intervalInt, err := strconv.Atoi(interval)
	if err != nil {
		log.Errorf("SYNC_TASKSTATUS_INTERVAL_SEC is not a number: %v", err)
		log.Infof("Set default value sec 3")

		intervalInt = 3
	}

	ticker := time.NewTicker(time.Duration(intervalInt) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			checkTasksStatus(e.TasksCache.GetCurrState())
		case <-ctx.Done():
			return
		}
	}
}

func checkTasksStatus(cache *map[string]e.Task, mu *sync.RWMutex) {
	mu.Lock()
	defer mu.Unlock()

	for id, task := range *cache {
		taskStatus, ok := task.(e.TaskStatus)
		if !ok {
			log.Warningf("Task %s is not a TaskStatus: WrongData", id)

			continue
		}

		if taskStatus.Completed {
			continue
		}

		creationTime, err := time.Parse(tools.TimeFormat, taskStatus.CreatedAt)
		if err != nil {
			log.Warningf("Task %s: Wrong Created Time: %v", id, err)

			continue
		}

		duration := time.Since(creationTime)
		taskStatus.Duration = duration.String()

		if taskStatus.WorkStatus != e.WIP && duration > 10*time.Second && duration < 6*time.Minute {
			taskStatus.WorkStatus = e.WIP
		}

		if duration >= 6*time.Minute {
			taskStatus.WorkStatus = e.DONE
			taskStatus.Completed = true
		}

		e.TasksCache.UpdateIntoWatcher(id, taskStatus)
	}

	log.Tracef("sync tasks duration done")
}
