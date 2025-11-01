package get

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	resp "go.mod/internal/lib/api/response"
	"go.mod/internal/service"
)

func New(log *slog.Logger) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.get.get.New"
		
		log := slog.With(
			slog.String("op", op),
			slog.String("requird_id", middleware.GetReqID(r.Context())),
		)
		
		city := chi.URLParam(r, "city")
		if city == "" {
			slog.Error("city is empty")
			render.JSON(w,r,resp.Error("not found"))
			return 	
		}

		res, err := service.GetWeather(city)
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			log.Error("failed to get city", "city", city)
			render.JSON(w,r,resp.Error("failed to get city"))
			return			
		}

		log.Info("got key", "temp", res)

		render.JSON(w, r, map[string]any{
    		"city":  city,
    		"temp": res,
		})
	}
}