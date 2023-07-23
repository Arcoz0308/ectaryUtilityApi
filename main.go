package main

import (
	"ectary/handlers/config"
	"ectary/handlers/mysql"
	"ectary/server"
	"ectary/utils/cron"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func main() {
	config.Init()
	mysql.Init()
	cron.LoadCron()
	app := fiber.New(
		fiber.Config{
			GETOnly: true,
		},
	)
	server.Load(app)
	err := app.Listen(":" + strconv.Itoa(config.C.Port))
	if err != nil {
		panic(err)
	}
}
