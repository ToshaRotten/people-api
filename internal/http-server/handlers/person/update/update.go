package update

import (
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"people-api/internal/http-server/utils/render"
	"people-api/internal/models/person"
	"people-api/internal/models/response"
)

type Request struct {
	Old person.Person `json:"old"`
	New person.Person `json:"new"`
}

type Response struct {
	response.Response
}

type PersonUpdater interface {
	UpdatePerson(new person.Person, old person.Person) error
}

func Update(log *slog.Logger, personUpdater PersonUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.Update"
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

		err = personUpdater.UpdatePerson(req.Old, req.New)
		if err != nil {
			log.Error("err:", err)
			//log.Info("person already exists", slog.String("name", req.Name))
			render.JSON(w, response.Error())
			return
		}

		log.Info("person updated")

		render.JSON(w, response.OK())
	}
}
