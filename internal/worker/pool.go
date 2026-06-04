package worker

import (
	"log"

	"github.com/RishavSinha20/cricstream/internal/analytics"
	"github.com/RishavSinha20/cricstream/internal/models"
)

type Pool struct {
	JobQueue chan models.MatchEvent
}

func NewPool(workerCount int) *Pool {

	pool := &Pool{
		JobQueue: make(chan models.MatchEvent, 1000),
	}

	for i := 0; i < workerCount; i++ {
		go pool.startWorker(i)
	}

	return pool
}

func (p *Pool) startWorker(id int) {

	for event := range p.JobQueue {

		analytics.ProcessEvent(event)

		log.Printf(
			"[Worker %d] Processed %s Score=%d",
			id,
			event.MatchID,
			event.Score,
		)
	}
}