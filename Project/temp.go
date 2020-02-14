package main

import (
	//"encoding/json"
	"gitlab.com/tolexo/aqua"
)

type Shipper struct {
	aqua.RestService `prefix:"catalog" root:"/shipper" version:"1"`
	createShipper    aqua.POST `url:"/"`
	updateShipper    aqua.PUT  `url:"/{shipper_id:[0-9]+}/"`
	listShipper      aqua.GET  `url:"/"`
	//getShipper       aqua.GET    `url:"/{shipper_id:[0-9]+}/"`
	//deleteShipper    aqua.DELETE `url:"/{shipper_id:[0-9]+}/"`
}

func (*Shipper) CreateShipper(req aqua.Aide) (httpCode int, data interface{}) {
	return 200, "Test string"
}

func (*Shipper) UpdateShipper(shipperID int, req aqua.Aide) (httpCode int, data interface{}) {
	return 200, "Test string"
}

type ListShipperResponseItem struct {
	ShipperId int           `json:"shipper_id"`
	Name      string        `json:"name"`
	Address   []ShipperAddr `json:"address"`
}
type ShipperAddr struct {
	AddressId int    `json:"Address_id"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
}

func (*Shipper) ListShipper(req aqua.Aide) (httpCode int, data interface{}) {
	Shipper_details := []ListShipperResponseItem{
		{
			ShipperId: 81,
			Name:      "first shipper",
			Address: []ShipperAddr{
				{
					AddressId: 93,
					Address:   "Shakti Nagar",
					City:      "Noida",
					State:     "UP",
				},
			},
		},
		{
			ShipperId: 82,
			Name:      "second shipper",
			Address: []ShipperAddr{
				{
					AddressId: 94,
					Address:   "Mayur Vihar",
					City:      "Delhi",
					State:     "Delhi",
				},
			},
		},
	}
	httpCode = 200
	data = Shipper_details
	return
}

func main() {
	server := aqua.NewRestServer()
	server.Port = 4305
	server.AddService(&Shipper{})
	server.Run()
}
