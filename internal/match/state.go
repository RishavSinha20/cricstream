package match

type MatchState struct {
	MatchID string

	Innings int

	Score   int
	Wickets int

	Over int
	Ball int

	Batsman string
	Bowler  string
}
