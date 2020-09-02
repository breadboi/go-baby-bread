package main

import (
	"github.com/Lukaesebrot/dgc"
)

// RegisterInfoGroup Registers all commands related to bot info
func RegisterInfoGroup(router *dgc.Router) {
	// Register the ping command
	router.RegisterCmd(&dgc.Command{
		Name:        "ping",
		Description: "Responds with 'pong!'",
		Usage:       "ping",
		Example:     "ping",
		IgnoreCase:  true,
		Handler:     PingCommand,
	})

	// Register the info command
	router.RegisterCmd(&dgc.Command{
		Name:        "info",
		Description: "Provides basic bot info",
		Usage:       "info",
		Example:     "info",
		IgnoreCase:  true,
		Handler:     InfoCommand,
	})

	// Register the support command
	router.RegisterCmd(&dgc.Command{
		Name:        "support",
		Description: "Provides a link to the support discord",
		Usage:       "support",
		Example:     "support",
		IgnoreCase:  true,
		Handler:     SupportCommand,
	})

	// Register the uptime command
	router.RegisterCmd(&dgc.Command{
		Name:        "uptime",
		Description: "Displays the current uptime of the bot",
		Usage:       "uptime",
		Example:     "uptime",
		IgnoreCase:  true,
		Handler:     UptimeCommand,
	})

	// Register the uptime command
	router.RegisterCmd(&dgc.Command{
		Name:        "invite",
		Description: "Provides a link to add me to your server",
		Usage:       "invite",
		Example:     "invite",
		IgnoreCase:  true,
		Handler:     InviteCommand,
	})
}

// RegisterSocialGroup Registers all commands related to bot social
func RegisterSocialGroup(router *dgc.Router) {
	// Register the w2g command
	router.RegisterCmd(&dgc.Command{
		Name:        "w2g",
		Description: "Generates a link to a private watch2gether room",
		Usage:       "w2g <Video Link>",
		Example:     "w2g https://www.youtube.com/watch?v=5qap5aO4i9A",
		IgnoreCase:  true,
		Handler:     Watch2getherCommand,
	})
}

// RegisterGamesGroup Registers all commands related to bot games
func RegisterGamesGroup(router *dgc.Router) {
	// Register the rndteams command
	router.RegisterCmd(&dgc.Command{
		Name:        "rndteams",
		Description: "Splits players in a voice channel into two random teams",
		Usage:       "rndteams <source voice> <target voice 1> <target voice 2>",
		Example:     "rndteams games blue orange",
		IgnoreCase:  true,
		Handler:     RandomizeTeams,
	})
}

// RegisterOwnerGroup Registers all owner related commands
func RegisterOwnerGroup(router *dgc.Router) {
	// Register the SetGame command
	router.RegisterCmd(&dgc.Command{
		Name:        "setgame",
		Description: "Sets the current game (requires owner)",
		Usage:       "setgame <game>",
		Example:     "setgame RocketLeague",
		IgnoreCase:  true,
		Handler:     SetGame,
	})

	// Register the Announcement command
	router.RegisterCmd(&dgc.Command{
		Name:        "announcement",
		Description: "Sends bot updates to all server owners (requires owner)",
		Usage:       "announcement <message>",
		Example:     "announcement This is an example message",
		IgnoreCase:  true,
		Handler:     Announcement,
	})
}