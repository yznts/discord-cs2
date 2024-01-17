package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var PracticeCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "practice",
		Description: "Practice mode",
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Compose sequence of commands
		commands := []string{
			"bot_quota_mode normal",
			"mp_autoteambalance 0",
			"bot_kick",
			"bot_add ct expert",
			"bot_add ct expert",
			"mp_restartgame 1",
		}
		// Compose sequence into one command
		command := strings.Join(commands, "; ")
		// Execute command
		res, err := RCON_DIAL.Execute(command)
		if err != nil {
			Respond(i.Interaction, "Error: "+err.Error())
		}
		// Respond with result
		Respond(i.Interaction, "Practice enabled \n\n"+res)
	},
}
