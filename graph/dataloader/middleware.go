package dataloader

import (
	"context"
	"net/http"

	"gorm.io/gorm"
)

type contextKey string

const loadersKey = contextKey("dataloaders")

// Middleware injects fresh dataloaders into the request context
func Middleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Create a brand new set of loaders for THIS specific request
			loaders := NewLoaders(db)

			// 2. Put them into the request context
			ctx := context.WithValue(r.Context(), loadersKey, loaders)

			// 3. Continue down the middleware chain with the updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// ForContext retrieves the loaders from context in your resolvers
func ForContext(ctx context.Context) *Loaders {
	loaders, _ := ctx.Value(loadersKey).(*Loaders)
	return loaders
}
