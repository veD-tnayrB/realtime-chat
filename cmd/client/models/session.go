package models

import "fmt"

var (
	ErrSessionAliasRequired = fmt.Errorf("An alias for the session is required.\n")
)

type Session struct {
	Alias        string
	ErrorChann   chan error
	ContactChann chan bool
	Contacts     map[string]*Contact
	CurrentChat  *Contact
}

func NewSession() *Session {
	errChann := make(chan error)
	contactChann := make(chan bool)
	contacts := map[string]*Contact{}
	return &Session{Alias: "An0n", Contacts: contacts, ErrorChann: errChann, ContactChann: contactChann}
}

func (s *Session) AddContact(host, alias string) error {
	if host == "" {
		return ErrContactHostRequired
	}

	if alias == "" {
		return ErrContactAliasRequired
	}

	if _, ok := s.Contacts[alias]; ok {
		s.CurrentChat = s.Contacts[alias]
		return nil
	}

	contact := NewContact(alias, host)
	err := contact.Connect(s.Alias)
	if err != nil {
		return err
	}

	go contact.StartListening(s.ErrorChann)
	s.Contacts[alias] = contact

	s.SetCurrentChat(contact)
	return nil
}

func (s *Session) SendMessage(content string) {
	if content == "" {
		return
	}

	err := s.CurrentChat.SendMessage(s.Alias, content)
	if err != nil {
		s.ErrorChann <- err
		return
	}
}

func (s *Session) SetCurrentChat(contact *Contact) {
	s.CurrentChat = contact
	go s.CurrentChat.ListenEvents(s.ContactChann, s.ErrorChann)
	s.ContactChann <- true
}

func (s *Session) Close() {
	close(s.ErrorChann)
	close(s.ContactChann)
}
