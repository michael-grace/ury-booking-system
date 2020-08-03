package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	"io/ioutil"
	"net/http"
)

func baseHTTPRequest(w http.ResponseWriter, r *http.Request) (config.HTTPRequest, error) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No Body")
		return config.HTTPRequest{}, err
	}

	var apiRequest config.HTTPRequest
	err = json.Unmarshal(reqBody, &apiRequest)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad JSON")
		return config.HTTPRequest{}, err
	}

	return apiRequest, nil
}
