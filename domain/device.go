package domain

type Device struct {
	Id          string `json:"id" validate:"required" dynamo:"Id"`
	DeviceModel string `json:"deviceModel" validate:"required" dynamo:"DeviceModel"`
	Name        string `json:"name" validate:"required" dynamo:"Name"`
	Note        string `json:"note" validate:"required" dynamo:"Note"`
	Serial      string `json:"serial" validate:"required" dynamo:"Serial"`
}
