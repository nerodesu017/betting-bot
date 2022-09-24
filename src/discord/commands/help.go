package commands

import "github.com/bwmarrin/discordgo"

var name = "help"

func HelpCommand(commands []*discordgo.ApplicationCommand) []*discordgo.ApplicationCommand {
	command := &discordgo.ApplicationCommand{
		Name:        name,
		Description: "Gives you a list of commands to use",
	}

	return append(commands, command)
}

func HelpHandler(commandHandlers map[string]func(discordClient *discordgo.Session, i *discordgo.InteractionCreate)) {
	commandHandler := func(discordClient *discordgo.Session, i *discordgo.InteractionCreate) {
		discordClient.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: `-
/help
/create-user
/get-info-user
/get-credits
/coin-flip
`,
			},
		})
	}
	commandHandlers[name] = commandHandler
}
