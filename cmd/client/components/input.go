package components

import (
	"log"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/veD-tnayrB/chat/cmd/client/models"
)

func Input(g *gocui.Gui, session *models.Session) error {
	maxX, maxY := g.Size()
	x0 := float32(maxX) / 5.5
	x1 := maxX - 1
	y1 := maxY - 1

	if v, err := g.SetView("input", int(x0), maxY-6, int(x1), y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Enter a message"
		v.Wrap = true
		v.Editable = true
		g.Highlight = true
		g.Cursor = true

	}
	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return sendMessage(g, v, session) }); err != nil {
		log.Panicln(err)
	}

	return nil

}

// sendMessage is called when Enter is pressed in the input view
func sendMessage(g *gocui.Gui, v *gocui.View, session *models.Session) error {
	chatView, err := g.View("chat")
	if err != nil {
		return err
	}

	inputText := strings.TrimSpace(v.Buffer())
	if inputText != "" {
		// fmt.Fprintln(chatView, fmt.Sprintf("[%s]: \n %v \n", session.Alias, inputText)) // Append input to chat view
		session.SendMessage(inputText)
	}

	v.Clear()
	v.SetCursor(0, 0)
	return keepDown(g, chatView)

}

func keepDown(_ *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	ox, _ := v.Origin()           // Get current horizontal origin (x)
	_, height := v.Size()         // Get height of the view (visible lines)
	lines := len(v.BufferLines()) // Get total number of lines in the view buffer
	maxScroll := lines - height   // Calculate max vertical scroll offset

	if maxScroll < 0 {
		maxScroll = 0 // Clamp to zero if content fits in view
	}

	v.SetOrigin(ox, maxScroll) // Scroll to bottom
	return nil
}
