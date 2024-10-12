package main

import (
	"github.com/krabiworld/lamarr/internal/api"
	"github.com/krabiworld/lamarr/internal/bot"
	"github.com/krabiworld/lamarr/internal/cfg"
	"github.com/krabiworld/lamarr/internal/db"
	"github.com/krabiworld/lamarr/internal/logger"
	repositoryImpl "github.com/krabiworld/lamarr/internal/repositories/impl"
	serviceImpl "github.com/krabiworld/lamarr/internal/services/impl"
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
