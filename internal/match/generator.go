package match

import (
	"fmt"
	"time"

	"github.com/RishavSinha20/cricstream/internal/models"
	"github.com/RishavSinha20/cricstream/internal/players"
	"github.com/RishavSinha20/cricstream/internal/utils"
)

func GenerateBallEvent(state *MatchState) models.MatchEvent {

	eventType := utils.RandomEventType()

	runs := utils.RandomRuns(eventType)

	wicket := false

	if eventType == "WICKET" {
		wicket = true
		state.Wickets++
	}

	state.Score += runs

	event := models.MatchEvent{
		EventID: fmt.Sprintf(
			"%s-%d",
			state.MatchID,
			time.Now().UnixNano(),
		),

		MatchID: state.MatchID,

		Innings: state.Innings,

		Over: state.Over,
		Ball: state.Ball,

		Batsman: players.Batsmen[state.Over%len(players.Batsmen)],
		Bowler:  players.Bowlers[state.Over%len(players.Bowlers)],

		Runs:   runs,
		Wicket: wicket,

		EventType: eventType,

		Score:   state.Score,
		Wickets: state.Wickets,

		Timestamp: time.Now().Unix(),
	}

	return event
}
