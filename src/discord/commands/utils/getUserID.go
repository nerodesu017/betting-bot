package utils

import "github.com/bwmarrin/discordgo"

func GetUserIDFromInteraction(i *discordgo.InteractionCreate) string {
	if i.Member == nil {
		return i.User.ID
	} else {
		return i.Member.User.ID
	}
}
