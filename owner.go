package main

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// SetGame Set the current game the bot is playing
func SetGame(dgoSession *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Message.Author.ID == "112068607864815616" {
		game := i.ApplicationCommandData().Options[0].StringValue()
		if game == "users" {
			guilds := dgoSession.State.Guilds

			// Inaccurate, implement a set to get an accurate user count
			userCount := 0
			for _, guild := range guilds {
				userCount += guild.MemberCount
			}

			guildCount := len(guilds)

			statusMessage := "Servers: " + strconv.Itoa(guildCount) + " | Users: " + strconv.Itoa(userCount)

			dgoSession.UpdateGameStatus(0, statusMessage)

		} else {
			dgoSession.UpdateGameStatus(0, game)
		}
	}
}

// Announcement Sends bot updates to all server owners
func Announcement(dgoSession *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Message.Author.ID == "112068607864815616" {
		message := i.ApplicationCommandData().Options[0].StringValue()
		guilds := dgoSession.State.Guilds

		builtResponse := strings.Builder{}
		builtResponse.WriteString("Error creating direct message channels with the following users")

		for _, guild := range guilds {
			ownerID := guild.OwnerID
			ownerDM, userError := dgoSession.UserChannelCreate(ownerID)

			if userError != nil {
				builtResponse.WriteString(ownerID)
			} else {
				_, sendError := dgoSession.ChannelMessageSend(ownerDM.ID, message)

				if sendError != nil {
					builtResponse.WriteString(ownerID)
				}
			}
		}

		if builtResponse.Len() > 1 {
			dgoSession.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: builtResponse.String(),
				},
			})
		}
	}
}
