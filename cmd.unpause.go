package main

import "github.com/bwmarrin/discordgo"

var UnpauseCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "unpause",
		Description: "Unpause game with 'mp_unpause_match'",
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Execute command
		res, err := RCON_DIAL.Execute("mp_unpause_match")
		if err != nil {
			Respond(i.Interaction, "Error: "+err.Error())
		}
		// Respond with result
		Respond(i.Interaction, "Game unpaused \n\n"+res)
	},
}
