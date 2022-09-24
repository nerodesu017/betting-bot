package coin

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	db "github.com/nerodesu017/discord-betting/src/database"
	"github.com/nerodesu017/discord-betting/src/discord/commands/games/config"
	"github.com/nerodesu017/discord-betting/src/discord/commands/utils"
	utils_ "github.com/nerodesu017/discord-betting/src/utils"
)

var name = "coin-flip"

func CoinFlipCommand(commands []*discordgo.ApplicationCommand) []*discordgo.ApplicationCommand {
	minVal := float64(config.MinBetValue)
	command := &discordgo.ApplicationCommand{
		Name:        name,
		Description: "Flip a coin, pays 19 to 10; Must bet a multiple of 10; You win if you get HEADS",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "bet",
				Description: "The amount you want to bet",
				Required:    true,
				MinValue:    &minVal,
				MaxValue:    float64(config.MaxBetValue),
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "heads-or-tails",
				Description: "Choose if you want 'heads' or 'tails'",
				Required:    true,
			},
		},
	}

	return append(commands, command)
}

func CoinFlipHandler(commandHandlers map[string]func(discordClient *discordgo.Session, i *discordgo.InteractionCreate)) {
	commandHandler := func(discordClient *discordgo.Session, i *discordgo.InteractionCreate) {

		userID := utils.GetUserIDFromInteraction(i)

		options := i.ApplicationCommandData().Options

		val := options[0].IntValue()
		choice := options[1].StringValue()

		if val%10 != 0 {
			discordClient.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Bet value must be a multiple of 10",
				},
			})
			return
		}

		if choice != "heads" && choice != "tails" {
			discordClient.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "You gotta choose 'heads' or 'tails'",
				},
			})
			return
		}

		var choiceNum int
		var otherChoice string
		if choice != "heads" {
			otherChoice = "heads"
			choiceNum = 1
		} else {
			otherChoice = "tails"
			choiceNum = 0
		}

		hasEnoughMoney := db.HasEnoughCredits(userID, val)

		if !hasEnoughMoney {
			discordClient.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "You don't have that many credits in your account",
				},
			})
		}

		randNum := utils_.RNG.Intn(2)
		if choiceNum != randNum {
			bal, _ := db.ChangeAndRetrieveBalance(userID, -val)
			discordClient.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "You lost..",
							Description: fmt.Sprintf("It's %s, you lost %d credits\nBalance: **%d** credits", otherChoice, val, bal),
						},
					},
				},
			})
			return
		} else {
			bal, _ := db.ChangeAndRetrieveBalance(userID, val/10*9)
			discordClient.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "You WON!",
							Description: fmt.Sprintf("It's %s, you won **%d** credits!\nBalance: **%d** credits", choice, val/10*9+val, bal),
						},
					},
				},
			})
			return
		}
	}
	commandHandlers[name] = commandHandler
}
