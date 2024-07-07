package models

import "time"

type OfferProfile struct {
	ID               uint
	Name             string
	LifeCycleStatus  string //Valid values are: new,approved,activated,deactivated
	ValidFor         OfferValidity
	ProductOfferings []ProductOffering
	Status           string //valid values: expiared,reserved,active, consumed
	AssignedDate     time.Time
}

type OfferValidity struct {
	StartDate   time.Time
	ExpiaryDate time.Time
	Duration    OfferDuration
}

type OfferDuration struct {
	ValueType   string //Valid Values: minutes,hours,days,weeks,months,years,[nil]
	Value       string
	ApplyPolicy string // valid Values: on-assignment, at-first-use
}

type Offer struct {
	Name             string
	ValidFor         ValidFor
	ProductOfferings []ProductOffering
	Status           string //valid values: expiared,reserved,active, consumed
	AssignedDate     time.Time
}

type ValidFor struct {
	StartDate time.Time
	EndDate   time.Time
}
