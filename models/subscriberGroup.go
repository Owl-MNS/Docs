package models

type SubscriberGroup struct {
	ID             uint
	Name           string
	Description    string
	Permissions    uint
	OrganizationID uint
}
