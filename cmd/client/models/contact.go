package models

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

var (
	ErrContactAliasRequired = fmt.Errorf("An alias for the contact is required")
	ErrContactHostRequired  = fmt.Errorf("A host for the contact is required")
	ErrConnectingToHost     = fmt.Errorf("Error connecting to the host")
	ErrGettingResponse      = fmt.Errorf("Error while trying to get message from host")
	ErrSubscribing          = fmt.Errorf("Error subscribing to contact events")
	ErrSendingMessage       = fmt.Errorf("Error sending message")
)

type Contact struct {
	Alias    string          `json:"alias"`
	Host     string          `json:"host"`
	events   chan Event      `json:"-"`
	Messages []Message       `json:"-"`
	conn     *websocket.Conn `json:"-"`
}

func NewContact(alias, host string) *Contact {
	messages := []Message{}
	return &Contact{Alias: alias, Host: host, Messages: messages}

}

func (c *Contact) Connect(sessionAlias string) error {
	if c.Alias == "" {
		return ErrContactAliasRequired
	}

	if c.Host == "" {
		return ErrContactHostRequired
	}

	conn, _, err := websocket.DefaultDialer.Dial(c.Host, nil)
	if err != nil {
		return ErrConnectingToHost
	}
	c.conn = conn

	selfContact := Contact{Alias: sessionAlias, Host: c.Host}
	parsed, err := json.Marshal(selfContact)
	if err != nil {
		return fmt.Errorf("%s: %s\n", ErrSubscribing, err.Error())
	}

	err = c.conn.WriteJSON(Event{Action: SubscribeEvent, Event: c.Alias, Data: parsed})
	if err != nil {
		return fmt.Errorf("%s: %s\n", ErrSubscribing, err.Error())
	}

	return nil
}

func (c *Contact) StartListening(errChan chan<- error) {
	defer func() {
		c.conn.Close()
	}()

	for {
		res := Response{}
		err := c.conn.ReadJSON(res)
		if err != nil {
			errChan <- ErrGettingResponse
			continue
		}

		if !res.Status {
			errChan <- res.Error
			continue
		}

		event := Event{}
		c.events <- event
	}

}

func (c *Contact) SendMessage(from string, content string) error {
	message := Message{Sender: from, Content: content}
	parsed, err := json.Marshal(message)
	if err != nil {
		return ErrSendingMessage
	}

	event := Event{Action: BroadcastEvent, Event: c.Alias, Data: parsed}
	err = c.conn.WriteJSON(event)
	if err != nil {
		return fmt.Errorf("%s: %s\n", ErrSendingMessage, err)
	}
	c.Messages = append(c.Messages, message)

	return nil
}

func (c *Contact) ListenEvents(contactChan chan<- bool, errChan chan<- error) {
	for {
		event := <-c.events
		if event.Action == BroadcastEvent {
			message := Message{}
			err := json.Unmarshal(event.Data, &message)
			if err != nil {
				errChan <- err
				continue
			}

			c.Messages = append(c.Messages, message)
			contactChan <- true
		}

	}
}
