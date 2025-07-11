package main

import (
	"github.com/krabiworld/lamarr/internal/api"
	"github.com/krabiworld/lamarr/internal/bot"
	"github.com/krabiworld/lamarr/internal/cfg"
	"github.com/krabiworld/lamarr/internal/db"
	"github.com/krabiworld/lamarr/internal/logger"
	repositoryImpl "github.com/krabiworld/lamarr/internal/repositories/impl"
	serviceImpl "github.com/krabiworld/lamarr/internal/services/impl"
	"github.com/krabiworld/lamarr/internal/uptime"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	uptime.Init()
	cfg.Init()
	logger.Init()
	conn := db.InitAndGet()

	guildRepository := repositoryImpl.NewGuildRepository(conn)
	guildService := serviceImpl.NewGuildServiceImpl(guildRepository)

	go api.Start()
	go bot.Start(guildService)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
