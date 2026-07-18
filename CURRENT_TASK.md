# Current Task - Scheduler Abstraction

## Status: ✅ COMPLETED

## Objective
Refactor goroutine usage to Scheduler abstraction.

## Rules
- Jangan ubah Chat API
- Jangan ubah AI Engine
- Jangan ubah Prompt Builder
- Jangan ubah Streaming
- Jangan buat Queue
- Jangan buat Redis Worker

## Scheduler Interface
**File:** `shared/scheduler/scheduler.go`
```go
type Scheduler interface {
    Schedule(ctx context.Context, job Job) error
}
```

## Job Interface
**File:** `shared/scheduler/job.go`
```go
type Job interface {
    Name() string
    Run(ctx context.Context) error
}
```

## Immediate Scheduler
**File:** `shared/scheduler/immediate_scheduler.go`
- Default implementation using goroutine
- ChatService tidak tahu goroutine digunakan

## Summary Job
**File:** `ai/summary/summary_job.go`
- SummaryJob implements Job interface
- Calls SummaryService.Summarize()

## Integration
**File:** `ai/engine/chat_service.go`
- Added scheduler field
- Changed direct goroutine call to scheduler.Schedule()

**File:** `bootstrap/router.go`
- Creates ImmediateScheduler
- Passes to NewChatService

## Future Compatibility
Architecture supports future implementations:
- WorkerScheduler
- RedisScheduler
- RabbitMQScheduler
- KafkaScheduler
- CloudTaskScheduler

## Validation
- ✅ go fmt ./...
- ✅ go vet ./...
- ✅ go build ./...

## Blocker
❌ Tidak ada
