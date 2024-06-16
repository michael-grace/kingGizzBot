package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
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
				Title  string `json:"title"`
				Artist string `json:"artist"`
			} `json:"track"`
		} `json:"nowPlaying"`
	} `json:"data"`
}

type songs struct {
	Songs map[string][]string `yaml:"songs"`
}

func emojiReplace(message, replacement string) string {
	letters := strings.Split(strings.ToLower(message), "")
	runes := []rune(strings.ToLower(message))
	for idx, r := range runes {
		if r >= 'a' && r <= 'z' {
			letters[idx] = fmt.Sprintf(":%s%s:", replacement, string(r))
		}
	}
	return strings.Join(letters, "")
}

func main() {

	var manual bool
	flag.BoolVar(&manual, "m", false, "manual run")

	var customMessage string
	flag.StringVar(&customMessage, "c", "", "custom message to send")

	var emoji string
	flag.StringVar(&emoji, "e", "", "emoji set to use")

	var development bool
	flag.BoolVar(&development, "d", false, "development mode")

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

	if botConfig.SlackHook == "" {
		botConfig.SlackHook = os.Getenv("SLACK_HOOK")
	}

	songCommentsFile, err := ioutil.ReadFile("songs.yml")

	if err != nil {
		panic(err)
	}

	var songComments songs
	err = yaml.Unmarshal(songCommentsFile, &songComments)

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
			comments, ok := songComments.Songs[songData.Data.NowPlaying.Track.Title]
			if ok {
				rand.Seed(time.Now().Unix())
				message = fmt.Sprintf("%s %s", botConfig.Message, comments[rand.Intn(len(comments))])
			}
		} else {
			message = customMessage
		}

		switch emoji {
		case "w":
			message = emojiReplace(message, "alphabet-white-")
		case "y":
			message = emojiReplace(message, "alphabet-yellow-")
		}

		if development {
			fmt.Println(message)
		} else {
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

}
