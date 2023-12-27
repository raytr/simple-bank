package b_log

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	TraceIDContextKey       = "Trace-ID"
	TraceIDRequestHeaderKey = "X-Correlation-ID"
)

func TraceIdentifierMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		traceId := request.Header.Get(TraceIDRequestHeaderKey)
		if traceId == "" {
			traceId = "req-" + uuid.NewString()
		}

		ctx := request.Context()
		ctx = context.WithValue(ctx, TraceIDContextKey, traceId)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

func TraceIdentifier(ctx context.Context, request *http.Request) context.Context {
	traceId := request.Context().Value(TraceIDContextKey)
	if traceId == "" {
		traceId = "req-" + uuid.NewString()
	}
	return context.WithValue(ctx, TraceIDContextKey, traceId)
}
