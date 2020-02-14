package model

import (
	"time"
)

type ShipperDB struct {
	tableName   struct{}  `sql:"Shipper"`
	ShipperID   int       `sql:"shipper_id,type:serial PRIMARY KEY"`
	Name        string    `sql:"name,type:varchar(50)"`
	Description string    `sql:"description,type:varchar(300)"`
	Mobile      string    `sql:"mobile,type:varchar(10)"`
	Email       string    `sql:"email,type:varchar(100)"`
	CreatedAt   time.Time `sql:"created_at,type:timestamp NOT NULL DEFAULT NOW()"`
	Status      string    `sql:"status,type:status_values NOT NULL DEFAULT 'enable'"`
}

type AddressDB struct {
	tableName struct{}  `sql:"Address"`
	AddressID int       `sql:"address_id,type:serial Primary Key"`
	City      string    `sql:"city,type:varchar(100)"`
	State     string    `sql:"state,type:varchar(100)"`
	Address   string    `sql:"address,type:varchar(300)"`
	CreatedAt time.Time `sql:"created_at,type:timestamp NOT NULL DEFAULT NOW()"`
	Status    string    `sql:"status,type:status_values NOT NULL DEFAULT 'enable'"`
}

type ShipperaddressDB struct {
	tableName        struct{}  `sql:"Shipper_Address"`
	ShipperAddressID int       `sql:"shipper_address,type:serial primary key"`
	FkShipperid      int       `sql:"fkshipper_id,type:int references Shipper(shipper_id)"`
	FkAddressid      int       `sql:"fkaddress_id,type:int references Address(address_id)"`
	AddressType      string    `sql:"address_type,type:address_values NOT NULL DEFAULT 'billing'"`
	CreatedAt        time.Time `sql:"created_at,type:timestamp NOT NULL DEFAULT NOW()"`
}
