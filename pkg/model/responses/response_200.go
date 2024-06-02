package responses

import "github.com/stainton/queue/pkg/model/resources"

type TaskResponse200 struct {

	// tasks waitting to be executed
	Waitting []resources.Task `json:"waitting"`

	// running tasks
	Running []resources.Task `json:"running"`
}

type ResultResponse200 struct {
	Results []resources.Result `json:"results"`
}
