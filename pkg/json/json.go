package json

import (
	gojson "encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Unmarshal(r io.Reader, v interface{}) error {
	return gojson.NewDecoder(r).Decode(&v)
}

func Marshal(r io.Reader, v interface{}) error {
	return gojson.NewDecoder(r).Decode(&v)
}

func Write(w http.ResponseWriter, code int, c interface{}) {
	b, err := gojson.MarshalIndent(c, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(b)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func Error(w http.ResponseWriter, code int, err error) {
	body := map[string]string{
		"code":    fmt.Sprint(code),
		"message": err.Error(),
	}
	Write(w, code, body)
}
