// Package scheduler provides job scheduling abstractions.
package scheduler

import "context"

// Job defines the interface for a schedulable job.
type Job interface {
	// Name returns the job name for logging.
	Name() string

	// Run executes the job.
	Run(ctx context.Context) error
}
