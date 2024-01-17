package main

import (
	"github.com/bwmarrin/discordgo"
)

var RconCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "rcon",
		Description: "Execute raw RCON command on server",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "command",
				Description: "RCON command to execute",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Extract args
		args := GetCommandArgs(i.ApplicationCommandData().Options)
		// Execute command
		res, err := RCON_DIAL.Execute(args["command"].StringValue())
		if err != nil {
			Respond(i.Interaction, "Error: "+err.Error())
		}
		// Respond with result
		Respond(i.Interaction, "RCON executed \n\n"+res)
	},
}
