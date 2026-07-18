// Package scheduler provides job scheduling abstractions.
package scheduler

import (
	"context"
	"log"
)

// ImmediateScheduler schedules jobs immediately using goroutines.
// This is the default scheduler implementation.
type ImmediateScheduler struct{}

// NewImmediateScheduler creates a new ImmediateScheduler.
func NewImmediateScheduler() *ImmediateScheduler {
	return &ImmediateScheduler{}
}

// Schedule executes the job asynchronously in a goroutine.
func (s *ImmediateScheduler) Schedule(ctx context.Context, job Job) error {
	go func() {
		// Use a new context since the parent may be cancelled
		jobCtx := context.Background()

		log.Printf("[scheduler] starting job: %s", job.Name())
		if err := job.Run(jobCtx); err != nil {
			log.Printf("[scheduler] job failed: %s, error: %v", job.Name(), err)
		} else {
			log.Printf("[scheduler] job completed: %s", job.Name())
		}
	}()
	return nil
}
