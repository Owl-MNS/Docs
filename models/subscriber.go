package models

type Subscriber struct {
	ID                uint
	UUID              string
	Details           SubscriberDetails
	Credentials       Credentials
	OrganizationID    uint
	SubscriberGroupID uint
	Offers            []Offer
}

type SubscriberDetails struct {
	ID           uint
	Name         string
	Email        string
	NationalID   string
	PassportID   string
	Mobile       string
	Phone        string
	SubscriberID uint
}

type Credentials struct {
	ID                  uint
	Username            string
	Password            string
	AuthenticationToken string
	SubscriberID        uint
}
