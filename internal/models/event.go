package models

type MatchEvent struct {
	EventID string `json:"event_id"`
	MatchID string `json:"match_id"`

	Innings int `json:"innings"`
	Over    int `json:"over"`
	Ball    int `json:"ball"`

	Batsman string `json:"batsman"`
	Bowler  string `json:"bowler"`

	Runs   int  `json:"runs"`
	Wicket bool `json:"wicket"`

	EventType string `json:"event_type"`

	Score   int `json:"score"`
	Wickets int `json:"wickets"`

	Timestamp int64 `json:"timestamp"`
}