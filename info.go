package main

import (
	"github.com/Lukaesebrot/dgc"
	"time"
)

var startTime time.Time

// Application start
func init() {
	startTime = time.Now()
}

// Logic for the ping command
func PingCommand(ctx *dgc.Ctx) {
	// Respond with the just set custom object
	ctx.RespondText("Pong")
}

// Logic for the info command
func InfoCommand(ctx *dgc.Ctx) {
	ctx.RespondText("Hello, I'm Baby Bread, my soul is golang :D")
}

// Logic for the support command
func SupportCommand(ctx *dgc.Ctx) {
	ctx.RespondText("discord.gg/UVpTZgS")
}

// Logic to get the uptime for the bot
func UptimeCommand(ctx *dgc.Ctx) {
	ctx.RespondText("I've been live for: " + time.Since(startTime).String())
}

// Logic to get the bot invite link
func InviteCommand(ctx *dgc.Ctx) {
	ctx.RespondText("To add Baby Bread to your server, follow this link.\nhttps://discordapp.com/oauth2/authorize?client_id=360277926966591488&scope=bot&permissions=536345815")
}
