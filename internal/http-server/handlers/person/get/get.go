package get

import (
	"log/slog"
	"net/http"
	"people-api/internal/http-server/utils/render"
	"people-api/internal/models/person"
	"people-api/internal/models/response"
)

type Request struct {
	Id int `json:"id"`
}

type Response struct {
	response.Response
	*person.Person
}

type PersonGetter interface {
	GetPerson(id int) (*person.Person, error)
}

func Get(log *slog.Logger, personGetter PersonGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.get.Get"
		log = log.With(slog.String("op", op))

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("invalid request", err)
			render.JSON(w, response.Error())
			return
		}

		log.Info("request accepted", slog.Any("request", req))

		p, err := personGetter.GetPerson(req.Id)
		if err != nil {
			log.Error("not found", err)
			render.JSON(w, response.Error())
			return
		}

		render.JSON(w, Response{
			Response: response.OK(),
			Person:   p,
		})

	}
}
