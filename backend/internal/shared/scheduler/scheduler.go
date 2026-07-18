// Package scheduler provides job scheduling abstractions.
package scheduler

import "context"

// Scheduler defines the interface for scheduling jobs.
type Scheduler interface {
	// Schedule schedules a job to be executed.
	Schedule(ctx context.Context, job Job) error
}
