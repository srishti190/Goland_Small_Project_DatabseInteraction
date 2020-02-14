package model

type ShipperReq struct {
	Name        string       `json:"Name" validate:"required"`
	Description string       `json:"Description" validate:"min=5,max=50"`
	Mobile      string       `json:"Mobile" validate:"required"`
	Email       string       `json:"Email" validate:"required"`
	SAddress    []AddressReq `json:"Address" validate:"dive"`
}

type AddressReq struct {
	City    string `json:"City" validate:"required"`
	State   string `json:"State" validate:"required"`
	Address string `json:"SAddress"`
}

type UpdateReq1 struct {
	ShipperID int    `json:"ShipperID" validate:"required"`
	Name      string `json:"Name" validate:"required"`
	Email     string `json:"Email" validate:"min=5,max=30|isdefault"`
	Mobile    string `json:"Mobile" validate:"len=10|isdefault"`
}
