package save

import (
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"people-api/internal/http-server/utils/render"
	"people-api/internal/models/person"
	"people-api/internal/models/response"
)

type Request struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

type Response struct {
	response.Response
}

type PersonSaver interface {
	SavePerson(person person.Person) error
}

func Save(log *slog.Logger, personSaver PersonSaver) http.HandlerFunc {
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

		p := person.Person{
			Name:       req.Name,
			Surname:    req.Surname,
			Patronymic: req.Patronymic,
		}

		err = personSaver.SavePerson(p)
		if err != nil {
			log.Error("err:", err)
			//log.Info("person already exists", slog.String("name", req.Name))
			render.JSON(w, response.Error())
			return
		}

		log.Info("person added")

		render.JSON(w, response.OK())
	}
}
