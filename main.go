package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

func main() {

	// Bot Startup
	token := LoadConfiguration("config.json").Secrets.Discord
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		panic(err)
	}

	err = session.Open()
	if err != nil {
		panic(err)
	}

	// Wait for the user to cancel the process
	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc
	}()

	// Create a dgc router
	router := dgc.Create(&dgc.Router{
		// We will allow '!' and 'example!' as the bot prefixes
		Prefixes: []string{
			"!",
		},
		IgnorePrefixCase: true,
		BotsAllowed:      false,
		Commands:         []*dgc.Command{},

		// We may inject our middlewares in here, but we will also use the corresponding method later on
		Middlewares: []dgc.Middleware{},

		// This handler gets called if the bot just got pinged (no argument provided)
		PingHandler: func(ctx *dgc.Ctx) {
			ctx.RespondText("Hello")
		},
	})

	// Register the default help command
	router.RegisterDefaultHelpCommand(session, nil)

	// Logging command usage
	router.RegisterMiddleware(func(next dgc.ExecutionHandler) dgc.ExecutionHandler {
		return func(ctx *dgc.Ctx) {
			messageContent := ctx.Event.Content
			guild, guildError := ctx.Session.Guild(ctx.Event.GuildID)

			log := messageContent

			if guildError == nil {
				log += (" - Executed in: " + guild.Name)
			}

			// Inject a custom object into the context
			ctx.CustomObjects.Set("Command", log)

			// You can retrieve the object like this
			obj := ctx.CustomObjects.MustGet("Command").(string)
			fmt.Println(obj)

			// Call the next execution handler
			next(ctx)
		}
	})

	RegisterInfoGroup(router)

	RegisterSocialGroup(router)

	RegisterGamesGroup(router)

	RegisterOwnerGroup(router)

	router.Initialize(session)
}
