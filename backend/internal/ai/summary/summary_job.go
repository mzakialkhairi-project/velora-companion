// Package summary provides conversation summary functionality.
package summary

import (
	"context"
	"log"
)

// SummaryJob wraps summary generation as a schedulable job.
type SummaryJob struct {
	summaryService *Service
	userID         uint64
	workspaceID    uint64
	conversationID uint64
}

// NewSummaryJob creates a new SummaryJob.
func NewSummaryJob(summaryService *Service, userID, workspaceID, conversationID uint64) *SummaryJob {
	return &SummaryJob{
		summaryService: summaryService,
		userID:         userID,
		workspaceID:    workspaceID,
		conversationID: conversationID,
	}
}

// Name returns the job name.
func (j *SummaryJob) Name() string {
	return "summary"
}

// Run executes the summary generation.
func (j *SummaryJob) Run(ctx context.Context) error {
	log.Printf("[summary-job] generating summary for conversation: %d", j.conversationID)
	return j.summaryService.Summarize(ctx, j.userID, j.workspaceID, j.conversationID)
}
