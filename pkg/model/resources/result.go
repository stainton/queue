package resources

type Result struct {

	// task name
	Name string `json:"name"`

	// success or fail
	Result bool `json:"result"`

	// cost time
	Cost int32 `json:"cost"`

	// log
	Log string `json:"log,omitempty"`
}
