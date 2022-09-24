package main

import (
	"os"
	"os/signal"

	_ "github.com/nerodesu017/discord-betting/src/config"
	db "github.com/nerodesu017/discord-betting/src/database"
	"github.com/nerodesu017/discord-betting/src/discord"
	"github.com/sirupsen/logrus"
)

func main() {
	defer db.Pool.Close()
	defer discord.DiscordClient.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	logrus.Info("Press Ctrl+C to exit")
	<-stop
}
