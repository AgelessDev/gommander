package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adjective-john/gommander"
	"github.com/bwmarrin/discordgo"

	"example/commands"
)

var (
	/*
	!! Use a .env file to store sensitive credentials instead in production.
	*/
	TOKEN = ""
	PREFIX = ""
)

func main() {
	s, err := discordgo.New("Bot " + TOKEN)

	if err != nil {
		log.Fatalf("Invalid bot params: %v", err)
		return
	}

	s.Identify.Intents = discordgo.IntentsGuildMessages

	handler := gommander.New(PREFIX)

	s.AddHandler(handler.DefaultHandler)

	err = s.Open()
	if err != nil {
		log.Fatalf("Error while connecting: %v", err)
		return
	}

	handler.RegisterOne("ping", "pong", commands.PingCommand)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	s.Close()
}