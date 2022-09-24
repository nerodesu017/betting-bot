package discord

import (
	"github.com/bwmarrin/discordgo"
	commands_package "github.com/nerodesu017/discord-betting/src/discord/commands"
	coin "github.com/nerodesu017/discord-betting/src/discord/commands/games/coins"
	users "github.com/nerodesu017/discord-betting/src/discord/commands/users"
)

var (
	commands        []*discordgo.ApplicationCommand                                                   = make([]*discordgo.ApplicationCommand, 0)
	commandHandlers map[string]func(discordClient *discordgo.Session, i *discordgo.InteractionCreate) = make(map[string]func(discordClient *discordgo.Session, i *discordgo.InteractionCreate))
)

func init() {
	commands = commands_package.HelpCommand(commands)
	commands_package.HelpHandler(commandHandlers)

	commands = users.CreateUserCommand(commands)
	users.CreateUserHandler(commandHandlers)

	commands = users.GetInfoUserCommand(commands)
	users.GetInfoUserHandler(commandHandlers)

	commands = users.GetCreditsCommand(commands)
	users.GetCreditsHandler(commandHandlers)

	commands = coin.CoinFlipCommand(commands)
	coin.CoinFlipHandler(commandHandlers)
}
