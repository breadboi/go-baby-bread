package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Lukaesebrot/dgc"
)

// Object representing part of the watch2gether json response
type Watch2getherRoom struct {
	Streamkey string `json:"streamkey"`
}

// Logic for the w2g command
func Watch2getherCommand(ctx *dgc.Ctx) {
	token := LoadConfiguration("config.json").Secrets.Watch2gether

	// Default url if one isn't provided
	url := "https://www.youtube.com/watch?v=DWcJFNfaw9c"

	if ctx.Arguments.Amount() > 0 {
		url = ctx.Arguments.Get(0).Raw()
	}

	// Create our request
	reqBody, marsherr := json.Marshal(map[string]string{
		"share":   url,
		"api_key": token,
	})

	if marsherr != nil {
		print(marsherr)
	}

	resp, posterr := http.Post("https://www.watch2gether.com/rooms/create.json",
		"application/json", bytes.NewBuffer(reqBody))

	if posterr != nil {
		print(posterr)
	}

	defer resp.Body.Close()

	body, ioerr := ioutil.ReadAll(resp.Body)

	if ioerr != nil {
		print(ioerr)
	}

	unmarshaled_body := Watch2getherRoom{}

	unmarshal_error := json.Unmarshal([]byte(body), &unmarshaled_body)

	if unmarshal_error != nil {
		print(unmarshal_error)
	}

	ctx.RespondText("https://www.watch2gether.com/rooms/" + unmarshaled_body.Streamkey)
}
