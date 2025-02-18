package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/binarstrike/chiya-chatbot-go/internal/command"
	"github.com/binarstrike/chiya-chatbot-go/internal/config"
	"github.com/binarstrike/chiya-chatbot-go/internal/event"
	dcgo "github.com/bwmarrin/discordgo"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	session, err := dcgo.New(cfg.Bot.Token)
	if err != nil {
		panic(err)
	}

	ev := event.NewEvent(session)
	ev.LoadAllEvents()

	err = session.Open()
	if err != nil {
		panic(err)
	}

	go command.DeploySlashCommand(session, cfg.Bot.AppId)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT)
	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGUSR2)
	log.Println("Press Ctrl+C to exit")
	<-stop

	session.Close()
}
