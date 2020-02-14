package main

import (
	"Project/shipper"
	"github.com/go-pg/pg"
	"gitlab.com/tolexo/aqua"
	"log"
	"os"
)

func main() {

	server := aqua.NewRestServer()
	server.Port = 4309
	server.AddService(&shipper.Shipper{})
	server.Run()
}
