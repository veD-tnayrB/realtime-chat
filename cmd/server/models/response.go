package models

type SuccResponse struct {
	Status bool        `json:"status"`
	Data   interface{} // Here implement a response
}

type ErrResponse struct {
	Status bool  `json:"status"`
	Error  error `json:"error"`
}
