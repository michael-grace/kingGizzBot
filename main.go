package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"net/http"
)

const (
	message            = "Hmmm, a King Gizzard song was played. Maybe we should VoNC Towells. :vonc: :towells:"
	slackHook          = ""
	nowPlayingEndpoint = ""
)

type slackMessage struct {
	Text string `json:"text"`
}

type nowPlaying struct {
	Data struct {
		NowPlaying struct {
			Track struct {
				Artist string `json:"artist"`
			} `json:"track"`
		} `json:"nowPlaying"`
	} `json:"data"`
}

func main() {

	var manual bool
	flag.BoolVar(&manual, "m", false, "manual run")
	flag.Parse()

	res, _ := http.Get(nowPlayingEndpoint)

	var songData nowPlaying

	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&songData)

	if err != nil {
		panic(err)
	}

	if manual || (len(songData.Data.NowPlaying.Track.Artist) >= 12 && songData.Data.NowPlaying.Track.Artist[:12] == "King Gizzard") {

		var toSend slackMessage = slackMessage{Text: message}
		bytesRepresentation, err := json.Marshal(toSend)
		if err != nil {
			panic(err)
		}
		_, err = http.Post(slackHook, "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			panic(err)
		}
	}

}
