package save_many

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
	Items []person.Person `json:"items"`
}

type Response struct {
	response.Response
}

type PersonSaver interface {
	SavePerson(person person.Person) error
}

func SaveMany(log *slog.Logger, personSaver PersonSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.SaveMany"
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

		for _, item := range req.Items {
			wg := sync.WaitGroup{}
			wg.Add(3)

			personToSave := person.Person{
				Name:       item.Name,
				Surname:    item.Surname,
				Patronymic: item.Patronymic,
			}

			go func() {
				personToSave.Age, err = get_age.GetAge(item.Name)
				if err != nil {
					return
				}
				wg.Done()
			}()

			go func() {
				personToSave.Sex, err = get_gender.GetGender(item.Name)
				if err != nil {
					return
				}
				wg.Done()
			}()

			go func() {
				nationalities, err := get_nationality.GetNationality(item.Name)
				personToSave.Nationality = nationalities[0]
				if err != nil {
					return
				}
				wg.Done()
			}()
			wg.Wait()

			err = personSaver.SavePerson(personToSave)
			if err != nil {
				log.Error("err:", err)
				render.JSON(w, response.Error())
				return
			}
		}
		log.Info("persons added")

		render.JSON(w, response.OK())
	}
}
