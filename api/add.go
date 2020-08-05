package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	"github.com/michael-grace/ury-booking-system/logic"
	"time"
	// "io/ioutil"
	"math/rand"
	"net/http"
)

// AddHandler takes the requests, and attempts to add them to the bookings
// and can also call logic functions if there are conflicts
func AddHandler(w http.ResponseWriter, r *http.Request, InProgressBookings map[int]config.InProgressBooking) {

	/*
	   First, sort the HTTP Request
	*/

	apiRequest, err := baseHTTPRequest("ADD", w, r)
	if err != nil {
		return
	}

	addRequest, ok := apiRequest.(config.BookingRequest)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad JSON - Add Object")
		return
	}

	/*
		Ask the DB and sort all conflicts
	*/

	var conflicts [][]config.Booking

	getConflictQuery := `SELECT * FROM bookings.bookings WHERE bookings.resource=$1 AND 
		NOT (bookings.end_time <= $2 OR bookings.start_time >= $3) 
		ORDER BY bookings.request_level, bookings.application_datetime DESC;`

	for _, requestTime := range addRequest.Requests {

		/*
			Each DB Query (essentially timeslot)
		*/

		rows, err := config.Database.Query(getConflictQuery, addRequest.Resource, requestTime.StartTime, requestTime.EndTime)
		if rows != nil {
			defer rows.Close()
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Database Query Failed %v \n", err)
			return
		}

		var individualBookingRequestConflicts []config.Booking

		/*
		   Each Conflict
		*/

		for rows.Next() {
			var conflictBooking config.Booking
			err = rows.Scan(
				&conflictBooking.BookingID,
				&conflictBooking.MemberID,
				&conflictBooking.RequestLevel,
				&conflictBooking.Resource,
				&conflictBooking.Preference,
				&conflictBooking.GivenResource,
				&conflictBooking.TimeslotID,
				&conflictBooking.StartTime,
				&conflictBooking.EndTime,
				&conflictBooking.PublicID,
				&conflictBooking.ApplicationDateTime,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "Database Output Flawed")
				return
			}
			individualBookingRequestConflicts = append(individualBookingRequestConflicts, conflictBooking)
		}

		conflicts = append(conflicts, individualBookingRequestConflicts)

	}

	returnToUser, err := logic.DealWithConflicts(addRequest, conflicts)

	if err != nil {
		return
	}

	userData := config.InProgressBooking{
		BookingRequest: addRequest,
		ManageType:     returnToUser,
	}

	for _, val := range userData.ManageType {
		if val.Header == logic.MANAGE {
			userData.ProgressID = rand.Intn(1000000000)
			InProgressBookings[userData.ProgressID] = userData
		}
	}

	if userData.ProgressID == 0 && userData.ManageType[0].Header == logic.ACCEPT {
		// Schedule These Bookings
		for _, b := range userData.ManageType {
			err = logic.TheBigScheduler(config.Booking{
				MemberID:            addRequest.MemberID,
				RequestLevel:        addRequest.RequestLevel,
				Resource:            addRequest.Resource,
				Preference:          addRequest.Preference,
				TimeslotID:          b.Booking.TimeslotID,
				StartTime:           b.Booking.StartTime,
				EndTime:             b.Booking.EndTime,
				ApplicationDateTime: time.Now(),
			})
			if err != nil {
				fmt.Fprintf(w, "Problem Scheduling %v \n", err)
			}
		}
	}

	jsonData, err := json.MarshalIndent(userData, "", "	")

	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))

}
