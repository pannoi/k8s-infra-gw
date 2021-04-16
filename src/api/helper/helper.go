package helper

import (
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
)

func WriteJSONResponse(w http.ResponseWriter, res interface{}) error {
	json, err := json.Marshal(res)

	if err != nil {
		return errors.Wrap(err, "Could not marshal json")
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

	return nil
}

func ParseBody(r io.Reader, out interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "Could not read body")
	}

	err = json.Unmarshal(b, out)
	if err != nil {
		return errors.Wrap(err, "Could not parse body")
	}

	return nil
}
