package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	"io/ioutil"
	"net/http"
)

// MoveHandler allows bookings to be moved, relying on the logic in Add
func MoveHandler(w http.ResponseWriter, r *http.Request) {

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
