package components

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/veD-tnayrB/chat/cmd/client/models"
)

func Chat(g *gocui.Gui, session *models.Session) error {
	maxX, maxY := g.Size()
	x0 := float32(maxX) / 5.5
	x1 := maxX - 1
	y1 := maxY - 7
	y0 := float32(maxY) / 12

	if v, err := g.SetView("chat", int(x0), int(y0), int(x1), y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Chat"
		v.Wrap = true
		g.Cursor = false
		setChatKeybinds(g)
		go listenChat(g, v, session)

	}

	return nil
}

func setChatKeybinds(g *gocui.Gui) {
	// Bind 'k' to scroll up in the "chat" view
	if err := g.SetKeybinding("chat", 'k', gocui.ModNone, scrollUp); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("chat", gocui.KeyCtrlU, gocui.ModNone, scrollUpFaster); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("chat", gocui.KeyCtrlB, gocui.ModNone, scrollUpFaster); err != nil {
		log.Panicln(err)
	}

	// Bind 'j' to scroll down in the "chat" view
	if err := g.SetKeybinding("chat", 'j', gocui.ModNone, scrollDown); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("chat", gocui.KeyCtrlD, gocui.ModNone, scrollDownFaster); err != nil {
		log.Panicln(err)
	}

}

func scrollUp(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	ox, oy := v.Origin()
	if oy > 0 {
		v.SetOrigin(ox, oy-1)
	}
	return nil
}

func scrollUpFaster(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	ox, oy := v.Origin()
	newOy := oy - 10
	if newOy < 0 {
		newOy = 0
	}
	v.SetOrigin(ox, newOy)
	return nil
}

func scrollDown(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	ox, oy := v.Origin()
	_, height := v.Size()
	// Get total lines in the view buffer
	lines := len(v.BufferLines())
	if oy < lines-height {
		v.SetOrigin(ox, oy+1)
	}
	return nil
}

func scrollDownFaster(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	ox, oy := v.Origin()
	_, height := v.Size()
	lines := len(v.BufferLines())
	maxScroll := lines - height
	if maxScroll < 0 {
		maxScroll = 0
	}
	newOy := oy + 10
	if newOy > maxScroll {
		newOy = maxScroll
	}
	v.SetOrigin(ox, newOy)
	return nil
}

func renderChat(v *gocui.View, session *models.Session) {
	v.Clear()

	for _, message := range session.CurrentChat.Messages {
		fmt.Printf("message: %s \n", message.Content)
		fmt.Fprintf(v, "\x1b[33m[%v]:\n \x1b[0m %s\n", message.Sender, message.Content)
	}
}

func listenChat(g *gocui.Gui, v *gocui.View, session *models.Session) {
	for {
		<-session.ContactChann
		if session.CurrentChat != nil {
			g.Update(func(g *gocui.Gui) error {
				renderChat(v, session)
				return nil
			})

		}

	}
}
