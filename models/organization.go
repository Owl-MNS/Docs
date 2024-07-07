package models

type Organization struct {
	ID                       uint
	UUID                     string
	Details                  OrganizationDetails
	Owner                    OrganizationOwner
	Balance                  float64
	AllowNagativeBalance     bool
	NegativeBalanceThreshold float64
}

type OrganizationDetails struct {
	ID             uint
	Name           string
	Address        string
	Email          string
	Mobile         string
	Phone          string
	OrganizationID uint
}

type OrganizationOwner struct {
	ID              uint
	Type            string //valid values are: legal, indivitual
	Name            string
	Address         string
	Email           string
	Mobile          string
	Phone           string
	LegalNationalID string
	OrganizationID  uint
}
