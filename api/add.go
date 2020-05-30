package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	"io/ioutil"
	"net/http"
)

// AddHandler takes the requests, and attempts to add them to the bookings
// and can also call logic functions if there are conflicts
func AddHandler(w http.ResponseWriter, r *http.Request) {

	var apiRequest config.HTTPRequest

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No Body")
	}

	err = json.Unmarshal(reqBody, &apiRequest)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad JSON")
	} else {
		addRequest, ok := apiRequest.Payload.(config.BookingRequest)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Bad JSON - Add Object")
		} else {
			var conflicts []config.Booking

			// TODO Logic Here
			getConflictQuery := "SELECT * FROM bookings.bookings WHERE bookings.resource_id=$1 AND "

			rows, err := config.Database.Query(getConflictQuery)
			defer rows.Close()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "Database Query Failed")
			} else {
				if rows.Next() {
					var conflictBooking config.Booking
					err = rows.Scan(&conflictBooking)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						fmt.Fprint(w, "Database Output Flawed")
					} else {
						conflicts = append(conflicts, conflictBooking)
						for rows.Next() {
							var conflictBooking config.Booking
							err = rows.Scan(&conflictBooking)
							if err != nil {
								w.WriteHeader(http.StatusInternalServerError)
								fmt.Fprint(w, "Database Output Flawed")
							} else {
								conflicts = append(conflicts, conflictBooking)
							}
						}
						logic.DealWithConflicts(addRequest, conflicts)
					}
				} else {
					// Schedule
				}
			}
		}
	}

}
