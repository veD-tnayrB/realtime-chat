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
		contactKeys := []string{}

		g.Cursor = false

		for key := range session.Contacts {
			contactKeys = append(contactKeys, key)
		}
		sort.Strings(contactKeys)

		setContactsKeybindings(g, contactKeys, &selectedIndex, session)
		renderContacts(v, &selectedIndex, contactKeys)
		go listenContacts(session, contactKeys)
	}

	return nil
}

func setContactsKeybindings(g *gocui.Gui, contactKeys []string, selectedIndex *int, session *models.Session) {
	g.SetKeybinding("contacts", 'j', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return cursorDown(v, contactKeys, selectedIndex) })
	g.SetKeybinding("contacts", 'k', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return cursorUp(v, contactKeys, selectedIndex) })
	g.SetKeybinding("contacts", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error { return selectContact(contactKeys, selectedIndex, session) })
}

func cursorDown(v *gocui.View, contactKeys []string, selectedIndex *int) error {
	if *selectedIndex < len(contactKeys)-1 {
		*selectedIndex++
	}
	renderContacts(v, selectedIndex, contactKeys)
	return nil
}

func cursorUp(v *gocui.View, contactKeys []string, selectedIndex *int) error {
	if *selectedIndex > 0 {
		*selectedIndex--
	}
	renderContacts(v, selectedIndex, contactKeys)
	return nil
}

func selectContact(contactsKeys []string, selectedIndex *int, session *models.Session) error {
	selected := contactsKeys[*selectedIndex]
	fmt.Printf("Selected contact: %v\n", selected)
	selectedContact := session.Contacts[selected]
	if selectedContact == nil {
		return fmt.Errorf("You just selected a contact that doesnt exists, tha fuck\n")
	}

	session.CurrentChat = selectedContact
	session.ChatChann <- true
	return nil
}

func renderContacts(v *gocui.View, selectedIndex *int, contactKeys []string) {
	v.Clear()

	for i, contact := range contactKeys {
		if i == *selectedIndex {
			// Highlight selected item, e.g., with ">" prefix or colors
			fmt.Fprintf(v, "> %s\n", contact)
		} else {
			fmt.Fprintf(v, "  %s\n", contact)
		}
	}
}

func listenContacts(session *models.Session, contactKeys []string) {
	for {
		<-session.ContactChann
		contactKeys = []string{}

		for key := range session.Contacts {
			contactKeys = append(contactKeys, key)
		}

		sort.Strings(contactKeys)
	}
}
