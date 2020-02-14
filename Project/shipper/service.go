package shipper

import (
	"Project/shipper/model"
	"Project/shipper/util"
	"encoding/json"
	"fmt"
	"gitlab.com/tolexo/aqua"
	//"gopkg.in/go-playground/validator.v9"
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
	var s model.ListShipperResponseItem
	req.LoadVars()
	//fmt.Println("here", req.Body)
	err := json.Unmarshal([]byte(req.Body), &s)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if err = util.V.Struct(s); err != nil {
			data = err.Error()
		}
	}
	httpCode = 200
	if err == nil {
		data = s
	}
	//fmt.Println(req.Body, httpCode)
	return
}

func (*Shipper) UpdateShipper(shipperID int, req aqua.Aide) (httpCode int, data interface{}) {
	return 200, "Test string"
}

func (*Shipper) ListShipper(req aqua.Aide) (httpCode int, data interface{}) {
	Shipper_details := []model.ListShipperResponseItem{
		{
			ShipperId: 81,
			Name:      "first shipper",
			Address: []model.ShipperAddr{
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
			Address: []model.ShipperAddr{
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
