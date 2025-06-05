package main

import (
	"github.com/veD-tnayrB/chat/cmd/client/components"
	"github.com/veD-tnayrB/chat/cmd/client/models"
	"github.com/veD-tnayrB/chat/cmd/client/views"
)

func main() {
	session := models.NewSession()
	view := views.View{Session: session}
	g := view.Init()

	defer func() {
		session.Close()
		g.Close()
	}()

	go components.ListenErrors(g, session.ErrorChann)

}
