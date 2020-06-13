package scales

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// BodyToJSON is ...
func BodyToJSON(w http.ResponseWriter, r *http.Request) (jsonBody ExcludePath, err error) {
	// Initialize
	length := 0

	//To allocate slice for request body
	length, err = strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse json
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
