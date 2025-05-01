package middlewares

import (
	"context"
	"net/http"

	"github.com/ZaphCode/F-SR-ChatApp/app"
	"github.com/google/uuid"
)

func Auth(next app.HandleFunc) app.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := app.GetSessionValue[string](r, "user_id")

		if err != nil || id == "" {
			return app.WriteJson(w, http.StatusUnauthorized, app.Response{
				Status: app.StatusFail, Msg: "Unauthorized", Error: "No session found",
			})
		}

		uid, err := uuid.Parse(id)

		if err != nil {
			return app.WriteJson(w, http.StatusInternalServerError, app.Response{
				Status: app.StatusFail, Msg: "Something went wrong", Error: err,
			})
		}

		ctx := context.WithValue(r.Context(), app.UserIDCtxKey, uid)

		return next(w, r.WithContext(ctx))
	}
}
