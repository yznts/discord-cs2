package main

import "github.com/bwmarrin/discordgo"

var RestartCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "restart",
		Description: "Restart game with 'mp_restartgame 1'",
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Execute command
		res, err := RCON_DIAL.Execute("mp_restartgame 1")
		if err != nil {
			Respond(i.Interaction, "Error: "+err.Error())
		}
		// Respond with result
		Respond(i.Interaction, "Game restarted\n\n"+res)
	},
}
