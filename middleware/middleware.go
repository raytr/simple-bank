package middleware

import (
	"net/http"

	"gibhub.com/raytr/simple-bank/helper/b_log"
)

type Middleware func(handler http.Handler) http.Handler

func Adapt(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func AddMiddleware(handler http.Handler) http.Handler {
	return Adapt(
		handler,
		CorsMiddleware,
		b_log.TraceIdentifierMiddleware,
	)
}
