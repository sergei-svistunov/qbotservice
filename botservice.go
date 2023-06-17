package botservice

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"

	"gopkg.qsoa.cloud/service"
)

type botService struct{}

var botImpl Bot

func init() {
	service.RegisterClientService(botService{})
}

func RegisterBot(bot Bot) {
	botImpl = bot
}

func (b botService) GetName() string { return "bgf-bot" }

func (b botService) Serve(_ net.Listener, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithCancel(context.Background())

	// Graceful shutdown on Interrupt signal
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt)
	go func() {
		<-sigC
		cancel()
	}()

	for ctx.Err() == nil {
		botImpl.StartGame(ctx)
	}
}
