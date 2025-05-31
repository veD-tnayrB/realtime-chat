package components

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func Error(g *gocui.Gui, message string) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("error", 0, 0, maxX-1, maxY/15); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Output"
		// v.Autoscroll = true
		v.FgColor = gocui.ColorRed
		fmt.Fprintln(v, message)
	}
	return nil
}

func showError(g *gocui.Gui, message string) error {
	view, err := g.SetCurrentView("error")
	if err != nil {
		return err
	}
	view.Clear()
	fmt.Fprintln(view, message)
	view.Highlight = true
	return nil
}

func closeError(g *gocui.Gui) error {
	view, err := g.SetCurrentView("error")
	if err != nil {
		return err
	}
	view.Clear()
	fmt.Fprintln(view, "")
	view.Highlight = true
	return nil
}
