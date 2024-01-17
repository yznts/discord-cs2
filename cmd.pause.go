package main

import "github.com/bwmarrin/discordgo"

var PauseCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "pause",
		Description: "Pause game with 'mp_pause_match'",
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Execute command
		res, err := RCON_DIAL.Execute("mp_pause_match")
		if err != nil {
			Respond(i.Interaction, "Error: "+err.Error())
		}
		// Respond with result
		Respond(i.Interaction, "Game paused \n\n"+res)
	},
}
