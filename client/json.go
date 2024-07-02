package client

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

func decodeJSONResponse[T any](body io.Reader) (T, error) {
	var res T

	if body == nil {
		return res, errors.New("no body to read")
	}

	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return res, errors.Wrap(err, "parse response")
	}

	return res, nil
}

func encodeJSONRequest(request any) (io.Reader, error) {
	encoded, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "marshal request")
	}
	return bytes.NewBuffer(encoded), nil
}
