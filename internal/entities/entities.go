package entities

const (
	WIP     = "WIP"     // Work in progress
	DONE    = "DONE"    // Done
	STARTED = "STARTED" // Started
)

type Task interface {
	plug()
}

type EntityInfo struct {
	ID string `json:"id" validate:"required"`
}

type TaskStatus struct {
	TaskInfo   *EntityInfo
	WorkStatus string
	CreatedAt  string
	Duration   string
	Completed  bool
}

func (t TaskStatus) plug() {}
