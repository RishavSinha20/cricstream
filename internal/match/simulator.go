package match

import (
	"time"

	"github.com/RishavSinha20/cricstream/internal/models"
)

func StartMatch(
	matchID string,
	events chan models.MatchEvent,
) {

	state := MatchState{
		MatchID: matchID,
		Innings: 1,
	}

	for {

		event := GenerateBallEvent(&state)

		events <- event

		state.Ball++

		if state.Ball == 6 {
			state.Over++
			state.Ball = 0
		}

		if state.Over == 20 {

			state.Over = 0
			state.Ball = 0
			state.Score = 0
			state.Wickets = 0
			state.Innings++

			if state.Innings > 2 {
				return
			}
		}

		time.Sleep(1 * time.Second)
	}
}