package models

type Response struct {
	Status bool
	Data   interface{} `json:"data,omitempty"`
	Error  error       `json:"error,omitempty"`
}
