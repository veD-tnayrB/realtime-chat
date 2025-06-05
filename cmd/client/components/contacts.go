package components

import (
	"fmt"
	"sort"

	"github.com/jroimartin/gocui"
	"github.com/veD-tnayrB/chat/cmd/client/models"
)

func ContactsSidebar(g *gocui.Gui, session *models.Session) error {
	selectedIndex := 0
	maxX, maxY := g.Size()
	y1 := maxY - 1
	y0 := float32(maxY) / 3.5
	if v, err := g.SetView("contacts", 0, int(y0), maxX/6, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Contacts"
		v.Wrap = true

		g.Cursor = false

		setContactsKeybindings(g, &selectedIndex, session)
		renderContacts(v, &selectedIndex, session)
		go listenContacts(g, v, &selectedIndex, session)
	}

	return nil
}

func setContactsKeybindings(g *gocui.Gui, selectedIndex *int, session *models.Session) {
	g.SetKeybinding("contacts", 'j', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return cursorDown(v, session, selectedIndex) })
	g.SetKeybinding("contacts", 'k', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return cursorUp(v, session, selectedIndex) })
	g.SetKeybinding("contacts", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return selectContact(g, selectedIndex, session) })
}

func cursorDown(v *gocui.View, session *models.Session, selectedIndex *int) error {
	contactKeys := getContactKeys(session)

	if *selectedIndex < len(contactKeys)-1 {
		*selectedIndex++
	}
	renderContacts(v, selectedIndex, session)
	return nil
}

func cursorUp(v *gocui.View, session *models.Session, selectedIndex *int) error {
	if *selectedIndex > 0 {
		*selectedIndex--
	}
	renderContacts(v, selectedIndex, session)
	return nil
}

func selectContact(g *gocui.Gui, selectedIndex *int, session *models.Session) error {
	contactKeys := getContactKeys(session)

	selected := contactKeys[*selectedIndex]
	selectedContact := session.Contacts[selected]
	if selectedContact == nil {
		return fmt.Errorf("You just selected a contact that doesnt exists, tha fuck\n")
	}

	session.SetCurrentChat(selectedContact)

	v, err := g.SetCurrentView("chat")
	if err != nil {
		return err
	}
	v.Title = selectedContact.Alias

	v, err = g.SetCurrentView("input")
	if err != nil {
		return err
	}

	v.Clear()
	return nil
}

func renderContacts(v *gocui.View, selectedIndex *int, session *models.Session) {
	v.Clear()
	contactKeys := getContactKeys(session)

	for i, contact := range contactKeys {
		if i == *selectedIndex {
			fmt.Fprintf(v, "> %s\n", contact)
		} else {
			fmt.Fprintf(v, "  %s\n", contact)
		}
	}
}

func listenContacts(g *gocui.Gui, v *gocui.View, selectedIndex *int, session *models.Session) {
	for {
		<-session.ContactChann
		g.Update(func(g *gocui.Gui) error {
			renderContacts(v, selectedIndex, session)
			return nil
		})
	}
}

func getContactKeys(session *models.Session) []string {
	contactKeys := []string{}
	for key := range session.Contacts {
		contactKeys = append(contactKeys, key)
	}
	sort.Strings(contactKeys)
	return contactKeys
}
