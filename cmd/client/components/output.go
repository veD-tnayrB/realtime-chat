package components

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func Error(g *gocui.Gui, message string) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("output", 0, 0, maxX-1, maxY/15); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Output"
		v.FgColor = gocui.ColorRed
		fmt.Fprintln(v, message)
	}
	return nil
}

func showOutput(g *gocui.Gui, message string) error {
	view, err := g.SetCurrentView("output")
	if err != nil {
		return err
	}
	view.Clear()
	fmt.Fprintln(view, message)
	view.Highlight = true
	return nil
}

func closeOutput(g *gocui.Gui) error {
	view, err := g.SetCurrentView("output")
	if err != nil {
		return err
	}
	view.Clear()
	fmt.Fprintln(view, "")
	view.Highlight = true
	return nil
}

func ListenErrors(g *gocui.Gui, errChann chan error) {
	for {
		err := <-errChann
		showOutput(g, err.Error())
	}

}
