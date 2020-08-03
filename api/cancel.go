package api

import (
	// "encoding/json"
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	// "io/ioutil"
	"net/http"
)

func cancelBookings(cancelRequest config.CancelRequest) []error {
	var errors []error
	for _, val := range cancelRequest.CancelBookingID {
		_, err := config.Database.Exec("DELETE FROM bookings.bookings WHERE booking_id = $1", val)
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors

}

// CancelHandler allows bookings to be removed from the system
func CancelHandler(w http.ResponseWriter, r *http.Request) {

	apiRequest, err := baseHTTPRequest("CNL", w, r)
	if err != nil {
		return
	}

	/*
	   Atttempts to put payload into cancel request
	   Goes through each booking id, and removing from DB
	   HTTP writes any errors that occur, or success
	*/

	cancelRequest, ok := apiRequest.(config.CancelRequest)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad JSON - Cancel Object")
	} else {

		errors := cancelBookings(cancelRequest)

		if len(errors) != 0 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Problem Removing Bookings", errors)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Bookings Removed")
		}

	}

}
