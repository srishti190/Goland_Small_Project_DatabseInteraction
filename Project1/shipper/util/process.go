package util

import (
	"Project1/shipper/model"
	"fmt"
	"github.com/go-pg/pg"
	"gitlab.com/tolexo/aqua"
)

var ShipperID int

func PCreateShipper(req aqua.Aide, db *pg.DB, shipperReq model.ShipperReq) (err error) {
	var (
		tx *pg.Tx
	)
	if tx, err = db.Begin(); err == nil {
		if ShipperID, err = InsertQueryForshipper(tx, shipperReq); err == nil {
			err = InsertQueryForaddress(tx, shipperReq.SAddress, ShipperID)
		}
		if err == nil {
			tx.Commit()
			fmt.Println("Process done")

		} else {
			tx.Rollback()
		}
	}
	return
}
