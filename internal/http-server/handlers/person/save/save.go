package save

import (
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"people-api/internal/http-client/handlers/get_age"
	"people-api/internal/http-client/handlers/get_gender"
	"people-api/internal/http-client/handlers/get_nationality"
	"people-api/internal/http-server/utils/render"
	"people-api/internal/models/person"
	"people-api/internal/models/response"
	"sync"
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
		const op = "handlers.url.Save"
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

		wg := sync.WaitGroup{}
		wg.Add(3)

		personToSave := person.Person{
			Name:       req.Name,
			Surname:    req.Surname,
			Patronymic: req.Patronymic,
		}

		go func() {
			personToSave.Age, err = get_age.GetAge(req.Name)
			if err != nil {
				return
			}
			wg.Done()
		}()

		go func() {
			personToSave.Sex, err = get_gender.GetGender(req.Name)
			if err != nil {
				return
			}
			wg.Done()
		}()

		go func() {
			nationalities, err := get_nationality.GetNationality(req.Name)
			personToSave.Nationality = nationalities[0]
			if err != nil {
				return
			}
			wg.Done()
		}()
		wg.Wait()

		if err != nil {
			log.Error("invalid request", err)
			render.JSON(w, response.Error())
			return
		}

		err = personSaver.SavePerson(personToSave)
		if err != nil {
			log.Error("err:", err)
			render.JSON(w, response.Error())
			return
		}

		log.Info("person added")

		render.JSON(w, response.OK())
	}
}
