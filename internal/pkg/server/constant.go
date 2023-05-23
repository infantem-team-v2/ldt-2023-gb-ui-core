package server

import (
	"gb-auth-gate/internal/auth/delivery/http"
	"gb-auth-gate/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
)

var (
	emptyHandlers = map[string]func(app *fiber.App) server.IHandler{
		"auth": http.NewAuthHandler,
	}
)
