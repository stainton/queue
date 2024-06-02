package resources

type Executor struct {

	// name of this executor
	Name string `json:"name,omitempty"`

	// ip of this executor
	Ip string `json:"ip,omitempty"`

	// executor's status
	Status string `json:"status,omitempty"`

	// task this executor is running
	Running string `json:"running,omitempty"`

	Config Config `json:"config,omitempty"`

	// tags of task this executor can run
	Tags []string `json:"tags,omitempty"`
}
