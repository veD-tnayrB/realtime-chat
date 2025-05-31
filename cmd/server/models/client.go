package models

import "github.com/gorilla/websocket"

type Client struct {
	Conn *websocket.Conn `json:"_"`
	Name string          `json:"name"`
}
