package model

/*type ListShipperResponseItem struct {
	ShipperId int           `json:"ShipperId"`
	Name      string        `json:"Name"`
	Address   []ShipperAddr `json:"Address" validate:"dive"`
}

type ShipperAddr struct {
	AddressId int    `json:"AddressId" validate:"required"`
	Address   string `json:"Address" validate:"required,min=1,max=50"`
	City      string `json:"City" validate:"eq=meerut"`
	State     string `json:"State" validate:"required,min=1,max=50"`
}*/
type Shipper struct {
	ShipperID   string `json:"shipper_id" sql:""`
	Name        string `json:"name" sql:""`
	Description string `json:"description" sql:""`
	Mobile      string `json:"mobile" sql:""`
	Email       string `json:"email" sql:""`
	CreatedAt   string `json:"created_at"`
	Status      string `json:"status"`
}

type Address struct {
	AddressID int    `json:"id"`
	City      string `json:"city"`
	State     string `json:"state"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
}

type Shipperaddress struct {
	ShipperAddress int    `json:"shipper_address" sql:""`
	FkShipperID    int    `json:"fkshipper_id" sql:""`
	FkAddressID    int    `json:"fkaddress_id" sql:""`
	AddressType    string `json:"address_type" sql:""`
	CreatedAt      string `json:"created_at" sql:""`
}

type Temp struct {
	ShipperID   string    `json:"shipper_id" sql:"shipper_id"`
	Name        string    `json:"name,omitempty" sql:"name"`
	Description string    `json:"description,omitempty"`
	Mobile      string    `json:"mobile,omitempty" sql:"mobile"`
	Email       string    `json:"email,omitempty"`
	Address     []Address `json:"address" sql:"address_info"`
}
