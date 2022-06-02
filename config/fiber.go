package config

import (
	"github.com/gofiber/fiber/v2"
	"skeleton-gofiber/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
