package main

import (
	"github.com/veD-tnayrB/chat/cmd/client/models"
	"github.com/veD-tnayrB/chat/cmd/client/views"
)

func main() {
	session := models.Session{Alias: "An0n", Contacts: map[string]*models.Contact{}}
	view := views.View{Session: &session}
	view.Init()
}
