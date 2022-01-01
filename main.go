package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", LoadConfiguration("config.json").Secrets.Discord, "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

var dgoSession *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	dgoSession, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "bot-info",
			Description: "Provides basic bot information",
		},
		{
			Name:        "bot-support",
			Description: "Provides a link to the support discord",
		},
		{
			Name:        "bot-uptime",
			Description: "Provides some stats on how long the bot has been online.",
		},
		{
			Name:        "w2g",
			Description: "Generates a link to a private watch2gether rooms.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "video-url",
					Description: "Provide a URL if you want a video preloaded into the room.",
					Required:    false,
				},
			},
		},
	}
	commandHandlers = map[string]func(dgoSession *discordgo.Session, i *discordgo.InteractionCreate){
		"bot-info":    InfoCommand,
		"bot-support": SupportCommand,
		"bot-uptime":  UptimeCommand,
		"w2g":         Watch2getherCommand,
	}
)

func init() {
	dgoSession.AddHandler(func(dgoSession *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(dgoSession, i)
		}
	})
}

func main() {
	dgoSession.AddHandler(func(dgoSession *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	err := dgoSession.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	for _, v := range commands {
		_, err := dgoSession.ApplicationCommandCreate(dgoSession.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	defer dgoSession.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")
}
