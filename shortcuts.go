package main

import "github.com/bwmarrin/discordgo"

func Respond(interaction *discordgo.Interaction, message string) {
	DISCORD.InteractionRespond(interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}
