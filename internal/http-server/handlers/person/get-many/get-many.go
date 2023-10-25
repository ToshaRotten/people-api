package get_many

import (
	"log/slog"
	"net/http"
	"people-api/internal/http-server/utils/pagination"
	"people-api/internal/http-server/utils/render"
	"people-api/internal/models/person"
	"people-api/internal/models/response"
	"strconv"
)

type Request struct{}

type Response struct {
	response.Response
	Items                 []person.Person `json:"items"`
	pagination.Pagination `json:"pagination"`
}

type PersonGetter interface {
	GetPersons(limit int, offset int) ([]person.Person, error)
	GetCountOfPersons() (int, error)
}

func GetMany(log *slog.Logger, personGetter PersonGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.getMany.GetMany"
		log = log.With(slog.String("op", op))

		page, err := strconv.Atoi(r.FormValue("page"))
		perPage, err := strconv.Atoi(r.FormValue("per_page"))
		if err != nil {
			log.Error("invalid request", err)
			render.JSON(w, response.Error())
			return
		}
		log.Info("request accepted")

		count, err := personGetter.GetCountOfPersons()
		if err != nil {
			log.Error("invalid request", err)
			render.JSON(w, response.Error())
			return
		}

		limit, offset := pagination.Calculate(page, perPage, count)

		persons, err := personGetter.GetPersons(limit, offset)
		if err != nil {
			log.Error("not found", err)
			render.JSON(w, response.Error())
			return
		}

		render.JSON(w, Response{
			Response:   response.OK(),
			Items:      persons,
			Pagination: pagination.Pagination{},
		})

	}
}
