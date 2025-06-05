package models

import "encoding/json"

var (
	SubscribeEvent   = "subscribe"
	UnsubscribeEvent = "unsubscribe"
	BroadcastEvent   = "broadcast"
)

type Event struct {
	Action string          `json:"action"`
	Event  string          `json:"event"`
	Data   json.RawMessage `json:"data"`
}
