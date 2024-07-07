package models

type ProductOffering struct {
	ID                    string
	Name                  string
	OfferingSpecification ProductOfferingSpecification
	Characteristic        ProductOfferingCharacteristic
	Price                 Price
	Tax                   Tax
	Discount              Discount
}

type ProductOfferingCharacteristic struct {
	ValueType string  `json:"value_type" example:"MB"`
	Value     float64 `json:"value" example:"100"`
}

type ProductOfferingSpecification struct {
	ID          string
	Name        string // Valid values are: Email, SMS, VoiceCall, PushNotification, WebHook, Telegram
	Type        string // valid values: product, service
	Description string
}

type Price struct {
	Value     float64
	ValueType string // Valid values: percentage, actual
}

type Tax struct {
	Value     float64
	ValueType string // Valid values: percentage, actual
}

type Discount struct {
	Value     float64
	ValueType string // Valid values: percentage, actual
}
