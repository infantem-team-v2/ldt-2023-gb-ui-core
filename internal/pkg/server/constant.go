package server

import (
	calculatorHttp "gb-ui-core/internal/calculator/delivery/http"
	"gb-ui-core/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
)

var (
	emptyHandlers = map[string]func(app *fiber.App) server.IHandler{
		"calc": calculatorHttp.NewCalculatorHandler,
	}
)
