package main

import (
	"encoding/json"
	"fmt"

	"github.com/RishavSinha20/cricstream/internal/match"
	"github.com/RishavSinha20/cricstream/internal/models"
	"github.com/RishavSinha20/cricstream/internal/utils"
)

func main() {

	utils.InitRandom()

	events := make(chan models.MatchEvent, 100)

	go match.StartMatch(
		"ipl_final_2026",
		events,
	)

	go match.StartMatch(
		"india_vs_australia",
		events,
	)

	go match.StartMatch(
		"rcb_vs_csk",
		events,
	)

	for event := range events {

		data, err := json.Marshal(event)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(string(data))
	}
}