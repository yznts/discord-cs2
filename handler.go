package main

import "github.com/bwmarrin/discordgo"

type Handler func(s *discordgo.Session, i *discordgo.InteractionCreate)
