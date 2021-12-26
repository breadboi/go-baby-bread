package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var startTime time.Time

// Application start
func init() {
	startTime = time.Now()
}

// InfoCommand Logic for the info command
func InfoCommand(dgoSession *discordgo.Session, i *discordgo.InteractionCreate) {
	dgoSession.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hi! I'm a discord bot programmed by Bread boi#0001. I was created using Go and discordgo.",
		},
	})
}

// SupportCommand Logic for the support command
func SupportCommand(dgoSession *discordgo.Session, i *discordgo.InteractionCreate) {
	dgoSession.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "discord.gg/UVpTZgS",
		},
	})
}

// UptimeCommand Logic to get the uptime for the bot
func UptimeCommand(dgoSession *discordgo.Session, i *discordgo.InteractionCreate) {
	dgoSession.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "I've been live for: " + time.Since(startTime).String(),
		},
	})
}
