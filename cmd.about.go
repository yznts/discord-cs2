package main

import "github.com/bwmarrin/discordgo"

// AboutCommand is an about command definition.
var AboutCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "about",
		Description: "Show information about bot",
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		DISCORD.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hello! I'm CS2 server manager, discord bot 🍷 \n" +
					"\n" +
					"I help CS2 admins to manage their servers with RCON. \n" +
					"I provide raw RCON interface as well as some useful shortcuts. \n" +
					"\n" +
					"Server address: " + SERV_ADDR + "\n" +
					"Server password: " + SERV_PASS + "\n" +
					"\n" +
					"To connect with CS2 console, use this command: \n" +
					"```connect " + SERV_ADDR + "; password " + RCON_PASS + "``` \n" +
					"",
			},
		})
	},
}
