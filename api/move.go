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

	/*
	   Moving Booking Requests
	   Tries to create a moveRequest
	   Then tries additions - TODO
	   Then removes stuff
	   API returns errors if they appear
	*/

	moveRequest, ok := apiRequest.Payload.(config.MoveRequest)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad JSON - Move Object")
	} else {

		// Additions First

		var bookingIDs []int
		for _, val := range moveRequest.MoveRequests {
			bookingIDs = append(bookingIDs, val.BookingID)
		}
		errors := cancelBookings(config.CancelRequest{CancelBookingID: bookingIDs})

		if len(errors) != 0 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Problem Moving Bookings", errors)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Bookings Moved")
		}
	}

}
