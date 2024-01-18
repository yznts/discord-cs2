package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"go.kyoto.codes/zen/v3/slice"
)

var (
	MODES = map[string][]int{
		"casual":         {0, 0},
		"competitive":    {0, 1},
		"wingman":        {0, 2},
		"weapons_expert": {0, 3},
		"deathmatch":     {1, 2},
		"demolition":     {1, 1},
	}
)

var ModeCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "mode",
		Description: "Change game mode",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "mode",
				Description: "Mode to change to",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices: slice.Map(
					[]string{
						"casual",
						"competitive",
						"wingman",
						"weapons_expert",
						"deathmatch",
						"demolition",
					},
					func(s string) *discordgo.ApplicationCommandOptionChoice {
						return &discordgo.ApplicationCommandOptionChoice{
							Name:  s,
							Value: s,
						}
					}),
				Required: true,
			},
		},
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Extract args
		args := GetCommandArgs(i.ApplicationCommandData().Options)
		// Extract mode pair
		mode, ok := MODES[args["mode"].StringValue()]
		if !ok {
			Respond(i.Interaction, "Error: Invalid mode. Contact bot owner.")
			return
		}
		// Execute command
		res, err := RCON_DIAL.Execute(fmt.Sprintf(
			"game_type %d; game_mode %d",
			mode[0], mode[1],
		))
		if err != nil {
			Respond(i.Interaction, "Error: "+err.Error())
		}
		// Respond with result
		Respond(i.Interaction, "Mode changed \n\n"+res)
	},
}
