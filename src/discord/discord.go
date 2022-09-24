package discord

import (
	"flag"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

var (
	BotToken = os.Getenv("DISCORD_SECRET_TOKEN")
)

var DiscordClient *discordgo.Session

func init() {
	var err error
	DiscordClient, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		logrus.Fatalf("Invalid bot parameters: %v", err)
	}

	DiscordClient.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	DiscordClient.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logrus.Infof("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = DiscordClient.Open()

	if err != nil {
		logrus.Fatalf("Cannot open the session: %v", err)
	}

	addCommands := flag.Bool("addCommands", false, "If it should add new commands to the discord bot")
	flag.Parse()

	if *addCommands {
		logrus.Infof("Adding %d commands...", len(commands))
		registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
		for i, v := range commands {
			cmd, err := DiscordClient.ApplicationCommandCreate(DiscordClient.State.User.ID, "", v)
			if err != nil {
				logrus.Fatalf("Cannot create '%v' command: %v", v.Name, err)
			}
			registeredCommands[i] = cmd
		}
		logrus.Info("Added commands successfully!")
	}

}
