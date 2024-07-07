package models

type PermissionSet struct {
	ID                     uint
	OrganizationLevelPerms OrganizationalLevelPerms
	AccessLevelPerms       AccessLevelPerms
	SubscriberLevelPerms   SubscriberLevelPerms
	PaymentLevelPerms      PaymentLevelPerms
	ReportLevelPerms       ReportLevelPerms
	OrganizationID         uint
}

type OrganizationalLevelPerms struct {
}

type AccessLevelPerms struct {
}

type SubscriberLevelPerms struct {
}

type PaymentLevelPerms struct {
}

type ReportLevelPerms struct {
}
