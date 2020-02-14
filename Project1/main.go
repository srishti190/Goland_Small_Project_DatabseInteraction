package main

import (
	"Project1/shipper"
	//"github.com/go-pg/pg"
	"gitlab.com/tolexo/aqua"
)

func main() {
	//defer db.Close()
	server := aqua.NewRestServer()
	server.Port = 4309
	server.AddService(&shipper.Shipper{})
	server.Run()
}
