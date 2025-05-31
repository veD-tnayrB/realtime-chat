package models

import (
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

var (
	ErrProviderHostRequired   = fmt.Errorf("A host is required to set a communication\n")
	ErrProviderContactingHost = fmt.Errorf("Error contacting the host \n")
)

type Provider struct {
	Host string
	conn *websocket.Conn
}

func (p *Provider) Start() error {
	if p.Host == "" {
		return ErrProviderHostRequired
	}

	serverURL := url.URL{Scheme: "ws", Host: p.Host, Path: "/ws"}
	// fmt.Printf("URL: %v\n", serverURL.String())
	conn, _, err := websocket.DefaultDialer.Dial(serverURL.String(), nil)
	if err != nil {
		return err
	}
	p.conn = conn
	go p.listen()

	return nil
}

func (p *Provider) listen() {
	defer p.conn.Close()

	for {
		res := Response{}
		err := p.conn.ReadJSON(res)
		if err != nil {
			return
		}

	}
}
