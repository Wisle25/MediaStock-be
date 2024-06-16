package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/wisle25/media-stock-be/commons"
	"github.com/wisle25/media-stock-be/infrastructures/server"
)

func main() {
	config := commons.LoadConfig(".")
	app := server.CreateServer(config)

	log.Fatal(
		app.Listen(fmt.Sprintf(":%s", config.ServerPort)),
	)
}
