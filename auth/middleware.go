package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/GigaDesk/eardrum-server/pkg/jwt"
	"github.com/rs/zerolog/log"
)

type favContextKey string

var k = favContextKey("user")

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			// Allow unauthenticated users in
			if header == "" {
				log.Info().Msg("Allowed unauthenticated user")
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			credentials, err := jwt.ParseToken(tokenStr)
			if err != nil {
				// put it in context
				ctx := context.WithValue(r.Context(), "error", err)
				// and call the next with our new context
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), k, credentials)
			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the token credentials from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) (*jwt.TokenCredentials, error){
	err := ctx.Value("error")
	if err!=nil{
		return nil, errors.New("Invalid Authentication Token")
	}
	user, _ := ctx.Value(k).(*jwt.TokenCredentials)
	return user, nil
}
