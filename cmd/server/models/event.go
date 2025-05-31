package models

import "fmt"

var (
	SubscribeEvent   = "suscribe"
	UnsubscribeEvent = "unsuscribe"
	BroadcastEvent   = "broadcast"
)

var (
	ActionRequired = fmt.Errorf("Action property must be defined. \n")
)

type Event struct {
	Action string
	Event  string
	Data   string
	From   *Client
}

func (e *Event) Validate() error {
	if e.Action == "" {
		return ActionRequired
	}

	return nil
}
