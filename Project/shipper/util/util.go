package util

import (
	"Project/shipper/model"
	"gopkg.in/go-playground/validator.v9"
)

var (
	V *validator.Validate
)

func init() {
	V = validator.New()
	V.RegisterStructValidation(func(arg validator.StructLevel) {
		p := arg.Current().FieldByName("Address")
		if p.Len() > 0 && p.Index(0).FieldByName("AddressId").Int() == arg.Current().FieldByName("ShipperId").Int() {
			arg.ReportError(arg.Current().FieldByName("ShipperId"), "ShipperId", "ShipperId", "validcheck", "")
		}
		return
	}, model.ListShipperResponseItem{})
}
