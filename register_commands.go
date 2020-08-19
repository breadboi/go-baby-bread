package main

import (
	"github.com/Lukaesebrot/dgc"
)

// Registers all commands related to bot info
func RegisterInfoGroup(router *dgc.Router) {
	// Register the ping command
	router.RegisterCmd(&dgc.Command{
		// We want to use 'obj' as the primary name of the command
		Name:        "ping",
		Description: "Responds with 'pong!'",
		Usage:       "ping",
		Example:     "ping",
		IgnoreCase:  true,
		Handler:     PingCommand,
	})

	// Register the info command
	router.RegisterCmd(&dgc.Command{
		// We want to use 'obj' as the primary name of the command
		Name:        "info",
		Description: "Provides basic bot info",
		Usage:       "info",
		Example:     "info",
		IgnoreCase:  true,
		Handler:     InfoCommand,
	})

	// Register the support command
	router.RegisterCmd(&dgc.Command{
		// We want to use 'obj' as the primary name of the command
		Name:        "support",
		Description: "Provides a link to the support discord",
		Usage:       "support",
		Example:     "support",
		IgnoreCase:  true,
		Handler:     SupportCommand,
	})

	// Register the uptime command
	router.RegisterCmd(&dgc.Command{
		// We want to use 'obj' as the primary name of the command
		Name:        "uptime",
		Description: "Displays the current uptime of the bot",
		Usage:       "uptime",
		Example:     "uptime",
		IgnoreCase:  true,
		Handler:     UptimeCommand,
	})

	// Register the uptime command
	router.RegisterCmd(&dgc.Command{
		// We want to use 'obj' as the primary name of the command
		Name:        "invite",
		Description: "Provides a link to add me to your server",
		Usage:       "invite",
		Example:     "invite",
		IgnoreCase:  true,
		Handler:     InviteCommand,
	})
}

// Registers all commands related to bot social
func RegisterSocialGroup(router *dgc.Router) {
	// Register the w2g command
	router.RegisterCmd(&dgc.Command{
		// We want to use 'obj' as the primary name of the command
		Name:        "w2g",
		Description: "Generates a link to a private watch2gether room",
		Usage:       "w2g <Video Link>",
		Example:     "w2g https://www.youtube.com/watch?v=5qap5aO4i9A",
		IgnoreCase:  true,
		Handler:     Watch2getherCommand,
	})
}