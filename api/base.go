package api

import (
	"github.com/michael-grace/ury-booking-system/config"
	"net/http"
)

func baseHTTPRequest(w http.ResponseWriter, r *http.Request)  config.HTTPRequest, error {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No Body")
		return nil, err
	}

	var apiRequest config.HTTPRequest
	err = json.Unmarshal(reqBody, &apiRequest)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad JSON")
		return nil, err
	}

	return apiRequest, nil
}
