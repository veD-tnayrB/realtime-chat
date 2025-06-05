package components

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/veD-tnayrB/chat/cmd/client/models"
)

func ContactAlias(g *gocui.Gui, session *models.Session) error {
	maxX, maxY := g.Size()
	x1 := maxX / 6
	y1 := float32(maxY) / 5
	y0 := float32(maxY) / 7

	if v, err := g.SetView("contact-alias", 0, int(y0), int(x1), int(y1)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "New contact alias"
		v.Wrap = true
		v.Editable = true
		g.Highlight = true

		if err := g.SetKeybinding("contact-alias", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return handleNewContactSave(g, v, session) }); err != nil {
			log.Panicln(err)
		}

	}
	return nil
}

func ContactHost(g *gocui.Gui, session *models.Session) error {
	maxX, maxY := g.Size()
	x1 := maxX / 6
	y1 := maxY / 4
	y0 := float32(maxY) / 4.5

	if v, err := g.SetView("contact-host", 0, int(y0), int(x1), int(y1)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "New contact host"
		v.Wrap = true
		v.Editable = true
		g.Highlight = true
		fmt.Fprintf(v, "ws://localhost:8080/ws")

		if err := g.SetKeybinding("contact-host", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return handleNewContactSave(g, v, session) }); err != nil {
			log.Panicln(err)
		}
	}
	return nil
}

func getInputText(g *gocui.Gui, viewName string) (string, error) {
	v, err := g.View(viewName)
	if err != nil {
		return "", err
	}
	input := v.Buffer()
	input = strings.TrimSpace(input)
	return input, nil
}

func handleNewContactSave(g *gocui.Gui, v *gocui.View, session *models.Session) error {
	inputText := strings.TrimSuffix(v.Buffer(), "\n")
	inputText = strings.TrimSpace(inputText)

	if inputText != "" {
		AddContact(g, session)
	}

	return nil

}

func AddContact(g *gocui.Gui, session *models.Session) {
	closeOutput(g)
	contactAlias, err := getInputText(g, "contact-alias")
	if err != nil {
		showOutput(g, "Something went wrong while trying to get the value of the contact-alias, please report this issie")
		g.SetCurrentView("contact-alias")
		return
	}

	contactHost, err := getInputText(g, "contact-host")
	if err != nil {
		showOutput(g, "Something went wrong while trying to get the value of the contact-host, please report this issie")
		g.SetCurrentView("contact-host")
		return
	}

	err = session.AddContact(contactHost, contactAlias)
	if err != nil {
		showOutput(g, err.Error())
		if err == models.ErrContactHostRequired {
			g.SetCurrentView("contact-host")
			return
		}
		if err == models.ErrContactAliasRequired {
			g.SetCurrentView("contact-alias")
			return
		}

		g.SetCurrentView("contact-host")
		return
	}

	closeOutput(g)
	showOutput(g, fmt.Sprintf(`Connection with "%s" established!`, contactAlias))

	v, _ := g.SetCurrentView("contact-alias")
	v.Clear()
	g.SetCurrentView("contacts")

}
