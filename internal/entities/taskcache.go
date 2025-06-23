package entities

import (
	"sync"
)

type taskCache struct {
	taskCache map[string]Task
	mu        *sync.RWMutex
}

var (
	// сделать структуру для кеша и объявить методы с мьютексами
	TasksCache = taskCache{
		taskCache: make(map[string]Task),
		mu:        new(sync.RWMutex),
	}
)

func (tC *taskCache) Set(key string, task Task) Task {
	tC.mu.Lock()
	defer tC.mu.Unlock()

	tC.taskCache[key] = task

	return task
}

func (tC *taskCache) Get(key string) Task {
	tC.mu.RLock()
	defer tC.mu.RUnlock()

	task := tC.taskCache[key]

	return task
}

func (tC *taskCache) Delete(key string) Task {
	tC.mu.Lock()
	defer tC.mu.Unlock()

	task := tC.taskCache[key]

	delete(tC.taskCache, key)

	return task
}

func (tC *taskCache) IsExists(key string) bool {
	tC.mu.RLock()
	defer tC.mu.RUnlock()

	_, ok := tC.taskCache[key]

	return ok
}

func (tC *taskCache) GetAll() map[string]Task {
	tC.mu.RLock()
	defer tC.mu.RUnlock()

	return tC.taskCache
}

func (tC *taskCache) GetCurrState() (*map[string]Task, *sync.RWMutex) {
	return &tC.taskCache, tC.mu
}

func (tC *taskCache) UpdateIntoWatcher(id string, task Task) { // deprecated 
	tC.taskCache[id] = task
}
