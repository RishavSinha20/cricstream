package analytics

type MatchStats struct {
	MatchID string

	Score   int
	Wickets int

	Balls int

	Fours int
	Sixes int

	RunRate float64
}
