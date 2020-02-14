package util

import (
	"Project1/shipper/model"
	"encoding/json"
	"fmt"
	"gitlab.com/tolexo/aqua"
	"gopkg.in/go-playground/validator.v9"
)

func VCreateShipper(req aqua.Aide) (shipperReq model.ShipperReq, err error) {
	req.LoadVars()
	if err = json.Unmarshal([]byte(req.Body), &shipperReq); err == nil {
		V := validator.New()
		if err = V.Struct(shipperReq); err == nil {
			fmt.Println("validate done")
		}
	}
	return
}
