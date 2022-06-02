package main

import (
	json2 "encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
	"skeleton-gofiber/config"
	"skeleton-gofiber/exception"
)

type Test struct {
	Name string
}

func main() {
	// Setup Configuration
	//configuration := config.New()
	//database := config.NewMongoDatabase(configuration)
	//database := config.NewMysqlDatabase(configuration)

	// Setup Repository

	// Setup Service

	// Setup Controller

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	test := Test{
		Name: "ara",
	}
	json, _ := json2.Marshal(test)
	logger.Info(string(json))

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
