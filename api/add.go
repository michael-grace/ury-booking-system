package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	"github.com/michael-grace/ury-booking-system/logic"
	"io/ioutil"
	"net/http"
)

// AddHandler takes the requests, and attempts to add them to the bookings
// and can also call logic functions if there are conflicts
func AddHandler(w http.ResponseWriter, r *http.Request) {

	/*
	   First, sort the HTTP Request
	*/

	apiRequest, err := baseHTTPRequest(w, r)
	if err != nil {
		return
	}

	addRequest, ok := apiRequest.Payload.(config.BookingRequest)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad JSON - Add Object")
		return
	}

	/*
		Ask the DB and sort all conflicts
	*/

	var conflicts [][]config.Booking

	// TODO Logic Here
	getConflictQuery := "SELECT * FROM bookings.bookings WHERE bookings.resource_id=$1 AND "

	for _, requestTime := range addRequest.Requests {

		/*
			Each DB Query (essentially timeslot)
		*/

		rows, err := config.Database.Query(getConflictQuery) // requestTime gets used here
		defer rows.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Database Query Failed")
			return
		}

		var individualBookingRequestConflicts []config.Booking

		/*
		   Each Conflict
		*/

		for rows.Next() {
			var conflictBooking config.Booking
			err = rows.Scan(&conflictBooking)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "Database Output Flawed")
				return
			}
			individualBookingRequestConflicts = append(individualBookingRequestConflicts, conflictBooking)
		}

		conflicts = append(conflicts, individualBookingRequestConflicts)

	}

	logic.DealWithConflicts(addRequest, conflicts)

}
