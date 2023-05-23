package main

import (
	"fmt"
	"gb-ui-core/internal/pkg/server"
)

// @title Backend-Driven-UI
// @description Service to provide UI specification frontend from backend
// @version 1.0.0
// @contact.name Docs developer
// @contact.url https://t.me/KlenoviySirop
// @contact.email KlenoviySir@yandex.ru

// @host ui.gb.ldt2023.infantem.tech
// @schemes https

func main() {
	if err := server.
		NewServer().
		MapHandlers().
		Run(); err != nil {
		panic(fmt.Sprintf("can't start server %+v", err))
	}
}
