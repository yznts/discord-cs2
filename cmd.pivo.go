package main

import "github.com/bwmarrin/discordgo"

var PivoCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "pivo",
		Description: "Send beer emoji",
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		Respond(i.Interaction, "🍺")
	},
}
