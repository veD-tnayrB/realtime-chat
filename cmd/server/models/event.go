package models

import (
	"encoding/json"
	"fmt"
)

var (
	SubscribeEvent   = "subscribe"
	UnsubscribeEvent = "unsubscribe"
	BroadcastEvent   = "broadcast"
)

var (
	ActionRequired = fmt.Errorf("Action property must be defined. \n")
)

type Event struct {
	Action string          `json:"action"`
	Event  string          `json:"event"`
	Data   json.RawMessage `json:"data"`
	From   *Client
}

func (e *Event) Validate() error {
	if e.Action == "" {
		return ActionRequired
	}

	return nil
}
