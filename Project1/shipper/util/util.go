package util

import (
	"Project1/shipper/model"
	"github.com/go-pg/pg"
)

func InsertQueryForshipper(tx *pg.Tx, x model.ShipperReq) (shipperID int, err error) {
	shipperDB := model.ShipperDB{
		Name:        x.Name,
		Email:       x.Email,
		Mobile:      x.Mobile,
		Description: x.Description,
	}
	if err = tx.Insert(&shipperDB); err == nil {
		shipperID = shipperDB.ShipperID
	}
	return
}

func InsertQueryForaddress(tx *pg.Tx, addr []model.AddressReq, shipperID int) (err error) {
	if len(addr) > 0 {
		var (
			addressDB        []model.AddressDB
			shipperAddressDB []model.ShipperaddressDB
		)
		for _, addrVal := range addr {
			addressDB = append(addressDB, model.AddressDB{
				City:    addrVal.City,
				State:   addrVal.State,
				Address: addrVal.Address,
			})
		}
		if err = tx.Insert(&addressDB); err == nil {
			for _, addrVal := range addressDB {
				shipperAddressDB = append(shipperAddressDB, model.ShipperaddressDB{
					FkShipperid: shipperID,
					FkAddressid: addrVal.AddressID,
					AddressType: "billing",
				})
			}
			err = tx.Insert(&shipperAddressDB)
		}
	}
	return
}
