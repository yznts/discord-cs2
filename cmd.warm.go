package main

import "github.com/bwmarrin/discordgo"

var WarmCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "warm",
		Description: "End warmup and start game with 'mp_warmup_end'",
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Execute command
		res, err := RCON_DIAL.Execute("mp_warmup_end")
		if err != nil {
			Respond(i.Interaction, "Error: "+err.Error())
		}
		// Respond with result
		Respond(i.Interaction, "Warmup ended \n\n"+res)
	},
}
