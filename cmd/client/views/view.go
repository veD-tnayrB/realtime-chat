package views

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/veD-tnayrB/chat/cmd/client/components"
	"github.com/veD-tnayrB/chat/cmd/client/models"
)

type View struct {
	gui     *gocui.Gui
	Session *models.Session
}

func (v *View) Init() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	v.gui = g

	g.SetManagerFunc(v.setLayout)
	g.Highlight = true
	g.SelFgColor = gocui.ColorYellow

	v.setKeybinding()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (v *View) setLayout(g *gocui.Gui) error {
	components.ContactsSidebar(g, v.Session)
	components.Chat(g, v.Session)
	components.Alias(g, v.Session)
	components.Input(g, v.Session)
	components.ContactAlias(g, v.Session)
	components.ContactHost(g, v.Session)
	components.Error(g, v.Session.Error)
	return nil
}
