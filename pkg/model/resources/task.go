package resources

// TaskFlags - flags for task execution
type TaskFlags struct {

	// retry regardless of result
	Retry int32 `json:"retry,omitempty"`

	// whether to print logs when the task succeeds
	Log bool `json:"log,omitempty"`

	Executor Executor `json:"executor"`

	// block executor before task execution
	BlockBefore bool `json:"block-before,omitempty"`

	// block executor after task execution
	BlockAfter bool `json:"block-after,omitempty"`
}

type Task struct {

	// task name, default value is used to configure the executor
	Name string `json:"name"`

	// priority of this task
	Priority int32 `json:"priority"`

	// block the executor, invalid if task name is not default value
	Block bool `json:"block,omitempty"`

	// skip all tasks, invalid if task name is not default value
	Finished bool `json:"finished,omitempty"`

	// skip tasks, invalid if task name is not default value
	Skips []string `json:"skips,omitempty"`

	// key word of this task
	Tags []string `json:"tags,omitempty"`

	Flags TaskFlags `json:"flags"`
}
