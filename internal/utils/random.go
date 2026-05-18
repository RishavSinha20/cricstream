package utils

import (
	"math/rand"
	"time"
)

var eventTypes = []string{
	"DOT",
	"SINGLE",
	"DOUBLE",
	"THREE",
	"FOUR",
	"SIX",
	"WICKET",
}

func InitRandom() {
	rand.Seed(time.Now().UnixNano())
}

func RandomEventType() string {
	return eventTypes[rand.Intn(len(eventTypes))]
}

func RandomRuns(eventType string) int {

	switch eventType {

	case "SINGLE":
		return 1

	case "DOUBLE":
		return 2

	case "THREE":
		return 3

	case "FOUR":
		return 4

	case "SIX":
		return 6

	default:
		return 0
	}
}
