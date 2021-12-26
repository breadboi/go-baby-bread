package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// Watch2getherRoom Object representing part of the watch2gether json response
type Watch2getherRoom struct {
	Streamkey string `json:"streamkey"`
}

// Watch2getherCommand Logic for the w2g command
func Watch2getherCommand(dgoSession *discordgo.Session, i *discordgo.InteractionCreate) {
	token := LoadConfiguration("config.json").Secrets.Watch2gether

	// Default url if one isn't provided
	url := "https://www.youtube.com/watch?v=DWcJFNfaw9c"

	if len(i.ApplicationCommandData().Options) > 0 {
		url = i.ApplicationCommandData().Options[0].StringValue()
	}

	// Create our request
	reqBody, marsherr := json.Marshal(map[string]string{
		"share":       url,
		"w2g_api_key": token,
	})

	if marsherr != nil {
		print(marsherr)
	}

	resp, posterr := http.Post("https://w2g.tv/rooms/create.json",
		"application/json", bytes.NewBuffer(reqBody))

	if posterr != nil {
		print(posterr)
	}

	defer resp.Body.Close()

	body, ioerr := ioutil.ReadAll(resp.Body)

	if ioerr != nil {
		print(ioerr)
	}

	unmarshaledBody := Watch2getherRoom{}

	unmarshalError := json.Unmarshal([]byte(body), &unmarshaledBody)

	if unmarshalError != nil {
		print(unmarshalError)
	}

	dgoSession.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "https://www.watch2gether.com/rooms/" + unmarshaledBody.Streamkey,
		},
	})
}
