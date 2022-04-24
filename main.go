package main

import (
	"fmt"

	"github.com/utiiz/go-server/config"
	"github.com/utiiz/go-server/model"
)

// func GetUsers(c *fiber.Ctx) error {
// 	users := []User{
// 		User{
// 			Username: "utiiz",
// 		},
// 		User{
// 			Username: "mavys",
// 		},
// 	}

// 	return json.NewEncoder(c).Encode(users)
// }

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println(config.Db.Address, err)
	}
	fmt.Println(config.Db.Address, err)
	// app := fiber.New()

	// app.Get("/users", GetUsers)

	// log.Fatal(app.Listen(":6666"))
	model.Setup(config)
}
