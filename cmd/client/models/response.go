package models

type Response struct {
	Status bool    `json:"Status"`
	Data   Message `json:"Data,omitempty"`
	Err    string  `json:"Err,omitempty"`
}
