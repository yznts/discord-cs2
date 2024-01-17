package main

import "github.com/bwmarrin/discordgo"

type Command struct {
	Command *discordgo.ApplicationCommand
	Handler
}

type CommandArgs map[string]*discordgo.ApplicationCommandInteractionDataOption

func GetCommandArgs(options []*discordgo.ApplicationCommandInteractionDataOption) CommandArgs {
	cmdmap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		cmdmap[opt.Name] = opt
	}
	return cmdmap
}
