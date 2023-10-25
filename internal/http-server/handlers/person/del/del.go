package del

import (
	"log/slog"
	"net/http"
	"people-api/internal/http-server/utils/render"
	"people-api/internal/models/response"
)

type Request struct {
	Id int64 `json:"id"`
}

type Response struct {
	response.Response
}

type PersonDeleter interface {
	DeletePerson(id int64) error
}

func Delete(log *slog.Logger, personDeleter PersonDeleter) http.HandlerFunc {
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

		err = personDeleter.DeletePerson(req.Id)
		if err != nil {
			log.Error("not found", err)
			render.JSON(w, response.Error())
			return
		}

		render.JSON(w, Response{
			Response: response.OK(),
		})

	}
}
