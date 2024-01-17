package main

import (
	"github.com/bwmarrin/discordgo"
	"go.kyoto.codes/zen/v3/slice"
)

var MapCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "map",
		Description: "Change map with 'changelevel <map>'",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "map",
				Description: "Map to change to",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices: slice.Map(
					[]string{
						"cs_italy",
						"cs_office",
						"de_ancient",
						"de_anubis",
						"de_dust2",
						"de_inferno",
						"de_mirage",
						"de_nuke",
						"de_overpass",
						"de_vertigo",
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
		// Execute command
		res, err := RCON_DIAL.Execute("changelevel " + args["map"].StringValue())
		if err != nil {
			Respond(i.Interaction, "Error: "+err.Error())
		}
		// Respond with result
		Respond(i.Interaction, "Map changed \n\n"+res)
	},
}
