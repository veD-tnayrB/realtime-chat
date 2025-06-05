package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Conns map[*Client]bool

type Hub struct {
	Conns         Conns
	Events        chan Event
	Mutx          sync.Mutex
	Subscriptions map[string]Conns
}

var (
	ErrSendingMessage = fmt.Errorf("Error sending message, message structure must be JSON")
)

func (h *Hub) Connect(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	client := Client{Conn: conn}

	h.Mutx.Lock()
	h.Conns[&client] = true
	h.Mutx.Unlock()

	go h.listen(&client)
}

func (h *Hub) listen(client *Client) {
	defer h.Disconnect(client)

	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			log.Printf("error getting message: %v \n", err)
			return
		}

		event := Event{}
		event.From = client
		err = json.Unmarshal(msg, &event)
		if err != nil {
			log.Printf("error parsing message to JSON: %v", err)
			return
		}

		h.Events <- event
	}
}

func (h *Hub) HandleEvents() {
	for {
		select {
		case event := <-h.Events:
			err := event.Validate()
			if err != nil {
				err := event.From.Conn.WriteJSON(ErrResponse{Status: false, Error: err})
				if err != nil {
					// Disconnects the client if the connection cant send the message
					h.Disconnect(event.From)
				}

				continue
			}

			fmt.Printf("%s: %s :\n", event.Action, event.Event)

			if event.Action == SubscribeEvent {
				h.Subscribe(&event)
				continue
			}

			if event.Action == UnsubscribeEvent {
				h.Unsubscribe(event.Event, event.From)
				continue
			}

			if event.Action == BroadcastEvent {
				h.Broadcast(&event)
				continue
			}
		}
	}
}

func (h *Hub) Subscribe(event *Event) {
	h.Mutx.Lock()
	defer h.Mutx.Unlock()

	if h.Subscriptions[event.Event] == nil {
		h.Subscriptions[event.Event] = make(Conns)
	}

	err := json.Unmarshal(event.Data, event.From)
	if err != nil {
		event.From.Conn.WriteJSON(ErrResponse{Status: true, Error: ErrSendingMessage})
		fmt.Printf("%s: %s", ErrSendingMessage, err)
		return
	}

	h.Subscriptions[event.Event][event.From] = true

	for client := range h.Subscriptions[event.Event] {
		client.Conn.WriteJSON(SuccResponse{Status: true, Data: event.From})
	}
}

func (h *Hub) Unsubscribe(event string, client *Client) {
	h.Mutx.Lock()
	defer h.Mutx.Unlock()

	eventMap := h.Subscriptions[event]
	if eventMap == nil {
		return
	}

	delete(eventMap, client)
}

func (h *Hub) Broadcast(event *Event) {
	h.Mutx.Lock()
	conns := h.Subscriptions[event.Event]
	h.Mutx.Unlock()
	if conns == nil {
		return
	}

	for conn := range conns {
		fmt.Printf("from: %s to: %s \n", event.From.Alias, conn.Alias)
		err := conn.Conn.WriteJSON(SuccResponse{Status: true, Data: event})
		if err != nil {
			// Close connection with the client if theres an issue sending the information through the conexion
			h.Disconnect(conn)
		}
	}
}

func (h *Hub) Disconnect(client *Client) {
	h.Mutx.Lock()
	delete(h.Conns, client)

	for _, conns := range h.Subscriptions {
		if conns == nil {
			continue
		}
		delete(conns, client)
	}

	h.Mutx.Unlock()
	client.Conn.Close()
}

func (h *Hub) Close() {
	close(h.Events)
}
