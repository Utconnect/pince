package responses

// swagger:model ResponseBody
type Body struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}
