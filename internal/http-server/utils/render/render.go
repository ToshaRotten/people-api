package render

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func DecodeJSON(r io.Reader, v any) error {
	const op = "utils.render.decodeJSON"
	b, err := io.ReadAll(r)
	if err != nil {
		return errors.New(op + "- failed to decode json")
	}

	err = json.Unmarshal(b, &v)
	if err != nil {
		return errors.New(op + "- failed to unmarshall json")
	}
	return nil
}

func JSON(w http.ResponseWriter, v any) {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(data)
	if err != nil {
		panic(err)
	}
}
