package middleware

import (
	"github.com/vardius/gorouter/v4"
	"net/http"

	"github.com/vardius/gocontainer"
	"placio-pkg/container"
)

// WithContainer wraps http.Handler with a container middleware
func WithContainer(requestContainer gocontainer.Container) gorouter.MiddlewareFunc {
	m := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := container.ContextWithContainer(r.Context(), requestContainer)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}

	return m
}
