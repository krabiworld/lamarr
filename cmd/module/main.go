package main

import (
	"module-go/internal/api"
	"module-go/internal/bot"
	"module-go/internal/cfg"
	"module-go/internal/db"
	"module-go/internal/logger"
	repositoryImpl "module-go/internal/repositories/impl"
	serviceImpl "module-go/internal/services/impl"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg.Init()
	logger.Init()
	conn := db.InitAndGet()

	guildRepository := repositoryImpl.NewGuildRepository(conn)
	guildService := serviceImpl.NewGuildServiceImpl(guildRepository)

	go api.Start()
	go bot.Start(guildService)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-c
}
