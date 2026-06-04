package analytics

import (
	"sync"

	"github.com/RishavSinha20/cricstream/internal/models"
)

var (
	statsMap = make(map[string]*MatchStats)
	mu       sync.RWMutex
)

func ProcessEvent(event models.MatchEvent) {

	mu.Lock()
	defer mu.Unlock()

	stats, exists := statsMap[event.MatchID]

	if !exists {

		stats = &MatchStats{
			MatchID: event.MatchID,
		}

		statsMap[event.MatchID] = stats
	}

	stats.Score = event.Score
	stats.Wickets = event.Wickets

	stats.Balls++

	switch event.EventType {

	case "FOUR":
		stats.Fours++

	case "SIX":
		stats.Sixes++
	}

	overs := float64(stats.Balls) / 6.0

	if overs > 0 {
		stats.RunRate = float64(stats.Score) / overs
	}
}

func GetStats(matchID string) (*MatchStats, bool) {

	mu.RLock()
	defer mu.RUnlock()

	stats, exists := statsMap[matchID]

	return stats, exists
}