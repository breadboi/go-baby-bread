package main

import (
	"strconv"

	"github.com/Lukaesebrot/dgc"
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

			ctx.Session.UpdateStatus(0, statusMessage)

		} else {
			ctx.Session.UpdateStatus(0, game)
		}
	}
}
