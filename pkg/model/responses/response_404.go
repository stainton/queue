package responses

type Response404 struct {
	Code int32 `json:"code"`

	Message string `json:"message"`
}
