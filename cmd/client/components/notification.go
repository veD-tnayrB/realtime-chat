package components

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
)

func showNotification(g *gocui.Gui, message string) {
	duration := 3 * time.Hour
	maxX, maxY := g.Size()
	g.Update(func(g *gocui.Gui) error {
		// Create or clear the notification view
		v, err := g.SetView("notification", maxX/4, maxY/4, maxX*3/4, maxY/4+2)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
		v.Clear()
		v.Frame = true
		v.Title = "Notification"
		fmt.Fprintln(v, message)
		return nil
	})

	// Hide notification after duration
	go func() {
		time.Sleep(duration)
		g.Update(func(g *gocui.Gui) error {
			g.DeleteView("notification")
			return nil
		})
	}()
}
