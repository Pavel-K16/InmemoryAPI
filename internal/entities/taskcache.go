package entities

import (
	"fmt"
	"sync"
	"taskapi/internal/logger"
	"time"
)

var (
	taskCache = make(map[string]Task)
	log       = logger.LoggerInit()
	mu        = new(sync.RWMutex)
)

func Create(info *EntityInfo) (Task, error) {
	mu.Lock()
	defer mu.Unlock()

	if info == nil || info.ID == "" {
		log.Errorf("Invalid task info")

		return nil, fmt.Errorf("invalid task info")
	}

	key := info.ID

	log.Tracef("Setting task %s in cache", key)

	if _, ok := taskCache[key]; ok {
		log.Warningf("Task %s already exists in cache", key)

		return nil, fmt.Errorf("task %s already exists in cache", key)
	}

	task := TaskStatus{
		TaskInfo:   info,
		WorkStatus: STARTED,
		CreatedAt:  time.Now(),
		Completed:  false,
		Duration:   "0m 0s",
	}

	// duration := end.Sub(start)

	// // Получаем минуты и секунды
	// minutes := int(duration.Minutes())
	// seconds := int(duration.Seconds()) % 60

	// fmt.Printf("Время выполнения: %d мин %d сек\n", minutes, seconds)

	taskCache[key] = task

	return &task, nil
}

func Get(info *EntityInfo) (*Task, error) {
	mu.RLock()
	defer mu.RUnlock()

	var task Task

	if info == nil || info.ID == "" {
		log.Errorf("Invalid task info")

		return nil, fmt.Errorf("invalid task info")
	}

	key := info.ID

	log.Tracef("Getting task %s from cache", key)

	if _, ok := taskCache[key]; !ok {
		log.Warningf("Task %s not found in cache", key)

		return nil, fmt.Errorf("task %s not found in cache", key)
	}

	task = taskCache[key]

	log.Tracef("Task %s added to cache", key)

	return &task, nil
}

func Delete(info *EntityInfo) (*Task, error) {
	mu.Lock()
	defer mu.Unlock()

	if info == nil || info.ID == "" {
		log.Errorf("Invalid task info")

		return nil, fmt.Errorf("invalid task info")
	}

	key := info.ID

	log.Tracef("Deleting task %s from cache", key)

	var task Task

	if _, ok := taskCache[key]; !ok {
		log.Warningf("Task %s not found in cache", key)

		return nil, fmt.Errorf("task %s not found in cache", key)
	}

	task = taskCache[key]

	delete(taskCache, key)

	log.Tracef("Task %s deleted from cache", key)

	return &task, nil
}

func GetAll() ([]Task, error) {
	log.Tracef("Getting all tasks from cache")

	mu.RLock()
	defer mu.RUnlock()

	var tasks []Task

	if len(taskCache) == 0 {
		log.Warningf("No tasks found in cache")

		return nil, fmt.Errorf("no tasks found in cache")
	}

	for _, task := range taskCache {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
