package main

import (
	"strconv"

	"github.com/lus/dgc"
)

// SetGame Set the current game the bot is playing
func SetGame(ctx *dgc.Ctx) {
	if ctx.Event.Author.ID == "112068607864815616" {
		game := ctx.Arguments.Get(0).Raw()
		if game == "users" {
			guilds := ctx.Session.State.Guilds

			// Inaccurate
			userCount := 0
			for _, guild := range guilds {
				userCount += guild.MemberCount
			}

			guildCount := len(guilds)

			statusMessage := "Servers: " + strconv.Itoa(guildCount) + " | Users: " + strconv.Itoa(userCount)

			ctx.Session.UpdateGameStatus(0, statusMessage)

		} else {
			ctx.Session.UpdateGameStatus(0, game)
		}
	}
}

// Announcement Sends bot updates to all server owners
func Announcement(ctx *dgc.Ctx) {
	if ctx.Event.Author.ID == "112068607864815616" {
		message := ctx.Arguments.AsSingle().Raw()
		guilds := ctx.Session.State.Guilds

		for _, guild := range guilds {
			ownerID := guild.OwnerID
			ownerDM, userError := ctx.Session.UserChannelCreate(ownerID)

			if userError != nil {
				ctx.RespondText("Error, a direct message could not be created for " + ownerID)
			} else {
				_, sendError := ctx.Session.ChannelMessageSend(ownerDM.ID, message)

				if sendError != nil {
					ctx.RespondText("Error, a direct message could not be sent to " + ownerID)
				}
			}
		}
	}
}
