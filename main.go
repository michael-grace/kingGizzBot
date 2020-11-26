package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

type config struct {
	Message            string `yaml:"message"`
	SlackHook          string `yaml:"slack"`
	NowPlayingEndpoint string `yaml:"nowPlaying"`
}

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

	var customMessage string
	flag.StringVar(&customMessage, "c", "", "custom message to send")

	flag.Parse()

	yamlFile, err := ioutil.ReadFile("config.yml")

	if err != nil {
		panic(err)
	}

	var botConfig config
	err = yaml.Unmarshal(yamlFile, &botConfig)

	if err != nil {
		panic(err)
	}

	res, _ := http.Get(botConfig.NowPlayingEndpoint)

	var songData nowPlaying

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&songData)

	if err != nil {
		panic(err)
	}

	if manual || customMessage != "" || (len(songData.Data.NowPlaying.Track.Artist) >= 12 && songData.Data.NowPlaying.Track.Artist[:12] == "King Gizzard") {

		var message string
		if customMessage == "" {
			message = botConfig.Message
		} else {
			message = customMessage
		}

		var toSend slackMessage = slackMessage{Text: message}
		bytesRepresentation, err := json.Marshal(toSend)
		if err != nil {
			panic(err)
		}
		_, err = http.Post(botConfig.SlackHook, "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			panic(err)
		}
	}

}
