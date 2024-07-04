package main

import (
	"module-go/internal/api"
	"module-go/internal/bot"
	"module-go/internal/cfg"
	"module-go/internal/db"
	"module-go/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg.Init()
	logger.Init()
	db.Init()

	go api.Start()
	go bot.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-c
}
