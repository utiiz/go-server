package main

import (
	"fmt"

	"github.com/utiiz/go-server/config"
	"github.com/utiiz/go-server/model"
	"github.com/utiiz/go-server/routes"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println(config.Db.Address, err)
	}
	fmt.Println(config.Db.Address, err)
	model.Setup(config)
	routes.Setup(config)
}
