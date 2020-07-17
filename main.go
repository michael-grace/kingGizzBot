package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type slackMessage struct {
	Text string `json:"text"`
}

type icecast struct {
	Icestats struct {
		Source []struct {
			Listenurl string `json:"listenurl"`
			Title     string `json:"title,omitempty"`
		} `json:"source"`
	} `json:"icestats"`
}

func VoncAlex() {
	var toSend slackMessage = slackMessage{Text: "Hmmm, a King Gizzard song was played. Maybe we should VoNC Towells. :vonc: :towells:"}
	bytesRepresentation, err := json.Marshal(toSend)
	if err != nil {
		panic(err)
	}
	_, err = http.Post("*****SLACK HOOK*****", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		panic(err)
	}

}

func main() {

	res, _ := http.Get("*****URY AUDIO STATUS URL*****")

	var songData icecast

	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&songData)

	if err != nil {
		panic(err)
	}

	for _, val := range songData.Icestats.Source {
		if val.Listenurl == "*****LISTEN LIVE URL*****" {
			artist := strings.Split(val.Title, " - ")[0]

			if artist == "King Gizzard and The Lizard Wizard" || artist == "King Gizzard And The Lizard Wizard" || artist == "King Gizzard & The Lizard Wizard" || artist == "King Gizzard and the Lizard Wizard & Mild High Club" {
				VoncAlex()
				break
			}

		}

	}

}
