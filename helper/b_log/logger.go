package b_log

import (
	"context"
	"os"
	"time"

	gokitlog "github.com/go-kit/log"
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
	Error(err error)
	WithContext(ctx context.Context) Logger
	Handle(ctx context.Context, err error)
	printf(logCollection *LogCollection)
}

type customLogger struct {
	Writer  gokitlog.Logger
	Context context.Context
	Caller  string
}

func NewLogger(caller string) Logger {
	return &customLogger{
		Writer: gokitlog.NewLogfmtLogger(os.Stderr),
		Caller: caller,
	}
}

func (l *customLogger) WithContext(ctx context.Context) Logger {
	lg := *l
	lg.Context = ctx
	return &lg
}

func (l *customLogger) Info(msg string) {
	l.printf(l.format(msg, InfoLevel))
}

func (l *customLogger) Error(err error) {
	l.printf(l.format(err.Error(), ErrorLevel))
}

// Handle - implement for ServerErrorHandler
func (l *customLogger) Handle(ctx context.Context, err error) {
	l.WithContext(ctx).Error(err)
}

func (l *customLogger) format(msg string, level string) *LogCollection {
	return &LogCollection{
		TraceID:  l.traceId(),
		DateTime: time.Now(),
		Level:    level,
		Message:  msg,
		Caller:   l.Caller,
	}
}

func (l *customLogger) traceId() string {
	var traceId string
	if l.Context != nil {
		if ctxTraceID, ok := l.Context.Value(TraceIDContextKey).(string); ok {
			traceId = ctxTraceID
		}
	}
	return traceId
}

func (l *customLogger) printf(logCollection *LogCollection) {
	_ = l.Writer.Log("trade-id", logCollection.TraceID, "level", logCollection.Level, "caller", logCollection.Caller, "message", logCollection.Message, "time", logCollection.DateTime)
}
