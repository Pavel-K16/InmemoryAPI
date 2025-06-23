package tasks

import (
	"fmt"
	e "taskapi/internal/entities"
	"taskapi/internal/logger"
	"taskapi/internal/tools"
	"time"
)

var (
	log = logger.LoggerInit()
)

func Create(info *e.EntityInfo) (e.Task, error) {

	if info == nil || info.ID == "" {
		log.Errorf("Invalid task info")

		return nil, fmt.Errorf("invalid task info")
	}

	key := info.ID

	log.Tracef("Setting task %s in cache", key)

	if e.TasksCache.IsExists(key) {
		log.Warningf("Task %s already exists in cache", key)

		return nil, fmt.Errorf("tas	k %s already exists in cache", key)
	}

	task := e.TaskStatus{
		TaskInfo:   info,
		WorkStatus: e.STARTED,
		CreatedAt:  time.Now().Format(tools.TimeFormat),
		Completed:  false,
		Duration:   "0m 0s",
	}

	// duration := end.Sub(start)

	// // Получаем минуты и секунды
	// minutes := int(duration.Minutes())
	// seconds := int(duration.Seconds()) % 60

	// fmt.Printf("Время выполнения: %d мин %d сек\n", minutes, seconds)

	e.TasksCache.Set(key, task)

	return &task, nil
}

func Get(info *e.EntityInfo) (*e.Task, error) {
	if info == nil || info.ID == "" {
		log.Errorf("Invalid task info")

		return nil, fmt.Errorf("invalid task info")
	}

	key := info.ID

	log.Tracef("Getting task %s from cache", key)

	if !e.TasksCache.IsExists(key) {
		log.Infof("Task %s not found in cache", key)

		return nil, fmt.Errorf("task %s not found in cache", key)
	}

	task := e.TasksCache.Get(key)

	log.Tracef("Task %s added to cache", key)

	return &task, nil
}

func Delete(info *e.EntityInfo) (*e.Task, error) {
	if info == nil || info.ID == "" {
		log.Errorf("Invalid task info")

		return nil, fmt.Errorf("invalid task info")
	}

	key := info.ID

	log.Tracef("Deleting task %s from cache", key)

	if !e.TasksCache.IsExists(key) {
		log.Infof("Task %s not found in cache", key)

		return nil, fmt.Errorf("task %s not found in cache", key)
	}

	task := e.TasksCache.Delete(key)

	log.Tracef("Task %s deleted from cache", key)

	return &task, nil
}

func GetAll() ([]e.Task, error) {
	log.Tracef("Getting all tasks from cache")

	var tasks []e.Task
	tC := e.TasksCache.GetAll()

	if len(tC) == 0 {
		log.Infof("No tasks found in cache")

		return nil, fmt.Errorf("no tasks found in cache")
	}

	for _, task := range tC {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
