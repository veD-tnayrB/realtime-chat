package views

import (
	"log"

	"github.com/jroimartin/gocui"
)

// Provides keybindings for general usage of the client
// Specific behaviors of individual components goes directly
// into the component.
func (v *View) setKeybinding() {
	if err := v.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, v.quit); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("input", gocui.KeyCtrlA, gocui.ModNone, v.focusContacts); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("input", gocui.KeyCtrlW, gocui.ModNone, v.focusChat); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("contacts", gocui.KeyCtrlW, gocui.ModNone, v.focusContactHost); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("contacts", gocui.KeyCtrlD, gocui.ModNone, v.focusChat); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("contacts", gocui.KeyCtrlS, gocui.ModNone, v.focusInput); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("alias", gocui.KeyCtrlS, gocui.ModNone, v.focusContactAlias); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("alias", gocui.KeyCtrlD, gocui.ModNone, v.focusChat); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("chat", gocui.KeyCtrlA, gocui.ModNone, v.focusContacts); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("chat", gocui.KeyCtrlS, gocui.ModNone, v.focusInput); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("contact-host", gocui.KeyCtrlW, gocui.ModNone, v.focusContactAlias); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("contact-host", gocui.KeyCtrlD, gocui.ModNone, v.focusChat); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("contact-host", gocui.KeyCtrlS, gocui.ModNone, v.focusContacts); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("contact-alias", gocui.KeyCtrlW, gocui.ModNone, v.focusAlias); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("contact-alias", gocui.KeyCtrlD, gocui.ModNone, v.focusChat); err != nil {
		log.Panicln(err)
	}

	if err := v.gui.SetKeybinding("contact-alias", gocui.KeyCtrlS, gocui.ModNone, v.focusContactHost); err != nil {
		log.Panicln(err)
	}

}

func (v *View) focusContacts(g *gocui.Gui, view *gocui.View) error {
	g.Cursor = false
	g.SetCurrentView("contacts")
	return nil
}

func (v *View) focusChat(g *gocui.Gui, view *gocui.View) error {
	g.Cursor = false
	g.SetCurrentView("chat")
	return nil
}

func (v *View) focusAlias(g *gocui.Gui, view *gocui.View) error {
	g.Cursor = true
	g.SetCurrentView("alias")
	return nil
}

func (v *View) focusContactAlias(g *gocui.Gui, view *gocui.View) error {
	g.Cursor = true
	g.SetCurrentView("contact-alias")
	return nil
}

func (v *View) focusContactHost(g *gocui.Gui, view *gocui.View) error {
	g.Cursor = true
	g.SetCurrentView("contact-host")
	return nil
}

func (v *View) focusInput(g *gocui.Gui, view *gocui.View) error {
	g.Cursor = true
	g.SetCurrentView("input")
	return nil
}

func (v *View) quit(g *gocui.Gui, view *gocui.View) error {
	return gocui.ErrQuit
}
