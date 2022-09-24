package users

import (
	"fmt"
	"time"
	"unicode"

	"github.com/bwmarrin/discordgo"
	db "github.com/nerodesu017/discord-betting/src/database"
	"github.com/nerodesu017/discord-betting/src/discord/commands/utils"
	"github.com/sirupsen/logrus"
)

var getInfoUserName = "get-info-user"

func GetInfoUserCommand(commands []*discordgo.ApplicationCommand) []*discordgo.ApplicationCommand {
	command := &discordgo.ApplicationCommand{
		Name:        getInfoUserName,
		Description: "Gives info about your account",
	}

	return append(commands, command)
}

func GetInfoUserHandler(commandHandlers map[string]func(discordClient *discordgo.Session, i *discordgo.InteractionCreate)) {
	commandHandlers[getInfoUserName] = func(discordClient *discordgo.Session, i *discordgo.InteractionCreate) {
		userID := utils.GetUserIDFromInteraction(i)

		user, err := db.SelectUser(userID)

		var content string
		if err != nil {
			content = "There was an error getting your info, make sure you have created an account first"
		} else {
			content = fmt.Sprintf(`-
Balance: **%d**
Last credit allocation: **%d** day(s) ago - (**%s** (YYYY-MM-DD))`, user.Balance, int(time.Since(user.LastAllocation).Hours()/24), user.LastAllocation.Local().Format("2006-01-02"))
		}

		discordClient.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
	}
}

var createUserName = "create-user"

func CreateUserCommand(commands []*discordgo.ApplicationCommand) []*discordgo.ApplicationCommand {
	command := &discordgo.ApplicationCommand{
		Name:        createUserName,
		Description: "Creates your account",
	}

	return append(commands, command)
}

func CreateUserHandler(commandHandlers map[string]func(discordClient *discordgo.Session, i *discordgo.InteractionCreate)) {
	commandHandlers[createUserName] = func(discordClient *discordgo.Session, i *discordgo.InteractionCreate) {
		userID := utils.GetUserIDFromInteraction(i)
		err := db.InsertUser(userID)
		var content string
		if err != nil {
			content = "Your account couldn't be created, contact the bot owner"
			logrus.Info(err)
		} else {
			content = "Account successfully created"
		}

		discordClient.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
	}
}

var getCreditsName = "get-credits"

func GetCreditsCommand(commands []*discordgo.ApplicationCommand) []*discordgo.ApplicationCommand {
	command := &discordgo.ApplicationCommand{
		Name:        getCreditsName,
		Description: "Gives you credits to play with (once per day)",
	}
	return append(commands, command)
}

func GetCreditsHandler(commandHandlers map[string]func(discordClient *discordgo.Session, i *discordgo.InteractionCreate)) {

	credits := 1000

	commandHandlers[getCreditsName] = func(discordClient *discordgo.Session, i *discordgo.InteractionCreate) {
		userID := utils.GetUserIDFromInteraction(i)
		err := db.CreditUser(userID, credits)
		var content string
		if err != nil {
			content = string(unicode.ToUpper(rune(err.Error()[0]))) + err.Error()[1:]
		} else {
			content = "Successfully credited your account"
		}

		discordClient.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
	}
}
