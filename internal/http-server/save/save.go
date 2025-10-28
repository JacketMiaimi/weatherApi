package save

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	resp "go.mod/internal/lib/api/response"
	"go.mod/internal/lib/logger/sl"
	red "go.mod/internal/redis"
)

type Request struct {
	City string `json:"city"`
	// Temp string `json:"temp"`
}

type Response struct {
	resp.Response
	Temp string `json:"temp"`
}

func New(log *slog.Logger, key string, temp string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.save.save.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")
			render.JSON(w, r, resp.Error("empty request"))
			return
		}

		if err != nil {
			log.Error("failed to decode request body", "error:",sl.Err(err))
			render.JSON(w,r,resp.Error("failed to decode request"))
			return 
		}

		err = red.SaveKey(key,temp)
		if err != nil {
			slog.Error("error to save key","error:",err)
		}

		
	}
}
