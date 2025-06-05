package models

import "github.com/gorilla/websocket"

type Client struct {
	Conn  *websocket.Conn `json:"_"`
	Alias string          `json:"alias"`
	Host  string          `json:"host"`
}
