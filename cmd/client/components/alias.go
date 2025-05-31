package components

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/veD-tnayrB/chat/cmd/client/models"
)

func Alias(g *gocui.Gui, session *models.Session) error {
	maxX, maxY := g.Size()
	x1 := maxX / 6
	y1 := float32(maxY) / 8
	y0 := float32(maxY) / 12

	if v, err := g.SetView("alias", 0, int(y0), int(x1), int(y1)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Your alias"
		v.Wrap = true
		v.Editable = true
		g.SetCurrentView("alias")
		g.Highlight = true
	}

	if err := g.SetKeybinding("alias", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return setAlias(g, v, session) }); err != nil {
		log.Panicln(err)
	}

	return nil

}

func setAlias(g *gocui.Gui, v *gocui.View, session *models.Session) error {
	inputText := strings.TrimSuffix(v.Buffer(), "\n")
	inputText = strings.TrimSpace(inputText)

	closeError(g)
	g.SetCurrentView("alias")
	v.Clear()
	fmt.Fprint(v, "") // Write inputText without extra newline

	if inputText != "" {
		session.Alias = inputText
		fmt.Fprint(v, inputText) // Write inputText without extra newline
		g.SetCurrentView("contact-alias")
		return nil
	}

	showError(g, models.ErrSessionAliasRequired.Error())
	g.SetCurrentView("alias")
	return nil
}
