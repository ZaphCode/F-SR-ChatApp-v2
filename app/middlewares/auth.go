package middlewares

import (
	"context"
	"net/http"

	"github.com/ZaphCode/F-SR-ChatApp/app"
	"github.com/google/uuid"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val, err := app.GetSessionValue(r, "user_id")

		if err != nil || val == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		id, ok := val.(uuid.UUID)

		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
