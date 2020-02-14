package model

type ListShipperResponseItem struct {
	ShipperId int           `json:"ShipperId"`
	Name      string        `json:"Name"`
	Address   []ShipperAddr `json:"Address" validate:"dive"`
}

type ShipperAddr struct {
	AddressId int    `json:"AddressId" validate:"required"`
	Address   string `json:"Address" validate:"required,min=1,max=50"`
	City      string `json:"City" validate:"eq=meerut"`
	State     string `json:"State" validate:"required,min=1,max=50"`
}
