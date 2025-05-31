package models

import "fmt"

var (
	ErrContactAliasRequired = fmt.Errorf("An alias for the contact is required,\n")
)

type Contact struct {
	Alias    string
	Provider *Provider
}
