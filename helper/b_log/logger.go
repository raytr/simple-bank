package b_log

import (
	"context"
	"time"

	gokit_log "github.com/go-kit/kit/log"
)

const (
	// InfoLevel is the default logging priority.
	InfoLevel = "info"

	// WarnLevel logs are more important than Info, but don't need individual human review.
	WarnLevel = "warn"

	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = "error"
)

type LogCollection struct {
	TraceID  string
	Level    string
	Caller   string
	Message  string
	DateTime time.Time
}

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	WithContext(ctx context.Context) Logger
}

type logger struct {
	Writer  gokit_log.Logger
	Context context.Context
}

func NewLogger(writer gokit_log.Logger) Logger {
	return &logger{
		Writer: writer,
	}
}

func (l *logger) WithContext(ctx context.Context) Logger {
	lg := *l
	lg.Context = ctx
	return &lg
}

func (l *logger) Info(msg string) {
	l.Writer.Log(l.format(msg, InfoLevel))
}

func (l *logger) Warn(msg string) {
	l.Writer.Log(l.format(msg, WarnLevel))
}

func (l *logger) Error(msg string) {
	l.Writer.Log(l.format(msg, ErrorLevel))
}

func (l *logger) format(msg string, level string) *LogCollection {
	return &LogCollection{
		DateTime: time.Now(),
		Level:    level,
		TraceID:  l.traceId(),
		Message:  msg,
	}
}

func (l *logger) traceId() string {
	var traceId string
	if l.Context != nil {
		if ctxTraceID, ok := l.Context.Value(TraceIDContextKey).(string); ok {
			traceId = ctxTraceID
		}
	}
	return traceId
}
