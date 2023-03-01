package main

import (
	"github.com/gerardva/go-api/config"
	"github.com/gerardva/go-api/database"
	"github.com/gerardva/go-api/server"
)

func main() {
	config.Init()
	database.Init()
	server.Init()
}