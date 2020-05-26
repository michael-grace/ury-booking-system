package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	"io/ioutil"
	"net/http"
)

// CancelHandler allows bookings to be removed from the system
func CancelHandler(w http.ResponseWriter, r *http.Request) {

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

	/*
	   Atttempts to put payload into cancel request
	   Goes through each booking id, and removing from DB
	   HTTP writes any errors that occur, or success
	*/

	cancelRequest, ok := apiRequest.Payload.(config.CancelRequest)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad JSON - Cancel Object")
	} else {
		var errors []error
		for _, val := range cancelRequest.CancelBookingID {
			_, err = config.Database.Exec("DELETE FROM bookings.bookings WHERE booking_id = $1", val)
			if err != nil {
				errors = append(errors, err)
			}
		}

		if len(errors) != 0 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Problem Removing Bookings", errors)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Bookings Removed")
		}
	}

}
