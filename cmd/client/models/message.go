package models

type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}
