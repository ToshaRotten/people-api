package save_many

import (
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"people-api/internal/http-server/utils/render"
	"people-api/internal/models/person"
	"people-api/internal/models/response"
)

type Request struct {
	Items []person.Person `json:"items"`
}

type Response struct {
	response.Response
}

type PersonSaver interface {
	SavePersons(persons []person.Person) error
}

func SaveMany(log *slog.Logger, personSaver PersonSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"
		log = log.With(slog.String("op", op))

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body: ", err)
			render.JSON(w, response.Error())
			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", err)
			render.JSON(w, response.Error())
			return
		}

		err = personSaver.SavePersons(req.Items)
		if err != nil {
			log.Error("err:", err)
			render.JSON(w, response.Error())
			return
		}

		log.Info("persons added")

		render.JSON(w, response.OK())
	}
}
