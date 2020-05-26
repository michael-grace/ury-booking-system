package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	"io/ioutil"
	"net/http"
)

// GetHandler deals with returning bookings within a timeframe in the body
func GetHandler(w http.ResponseWriter, r *http.Request) {

	var apiRequest config.HTTPRequest

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No Body")
	}

	err = json.Unmarshal(reqBody, &apiRequest)

	if err != nil {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Bad JSON")
		}
	}

}
