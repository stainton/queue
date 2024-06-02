package responses

type Response500 struct {
	// error code
	Code int32 `json:"code"`

	// error information
	Message string `json:"message"`
}
