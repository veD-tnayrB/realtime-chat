package models

import "fmt"

var (
	ErrSessionAliasRequired = fmt.Errorf("An alias for the session is required.\n")
)

type Session struct {
	Error        string
	Alias        string
	ContactChann chan bool
	ChatChann    chan bool
	Contacts     map[string]*Contact
	CurrentChat  *Contact
}

func (s *Session) AddContact(host, alias string) error {
	if host == "" {
		return ErrProviderHostRequired
	}

	if alias == "" {
		return ErrContactAliasRequired
	}

	if _, ok := s.Contacts[alias]; ok {
		s.CurrentChat = s.Contacts[alias]
		return nil
	}

	fmt.Printf("host: %v\n", host)

	provider := Provider{Host: host}
	err := provider.Start()
	if err != nil {
		return fmt.Errorf("%v: %v\n", ErrProviderContactingHost.Error(), err)
	}
	contact := Contact{Alias: alias, Provider: &provider}
	s.Contacts[alias] = &contact

	return nil
}
