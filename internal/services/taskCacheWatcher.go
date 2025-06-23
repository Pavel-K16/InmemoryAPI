package services

import (
	"context"
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

	ticker := time.NewTicker(3 * time.Second)
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

}
