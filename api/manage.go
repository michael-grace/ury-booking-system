package api

import (
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	"github.com/michael-grace/ury-booking-system/logic"
	"net/http"
	"time"
)

// ManageHandler is for dealing with user confirmations
func ManageHandler(w http.ResponseWriter, r *http.Request, inProgressBookings map[int]config.InProgressBooking) {

	/*
		First, sort the HTTP request
	*/

	apiRequest, err := baseHTTPRequest(w, r)
	if err != nil {
		return
	}

	manageRequest, ok := apiRequest.Payload.(config.ManageRequest)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad JSON - Manage Object")
		return
	}

	thisBooking, ok := inProgressBookings[manageRequest.ProgressID]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No Booking with given ID. Possible Timeout")
		return
	}

	for key, val := range thisBooking.ManageType {
		if val.Header == logic.ACCEPT && manageRequest.UserResponses[key] == true {
			// Schedule This Booking
			logic.TheBigScheduler(config.Booking{
				MemberID:            thisBooking.BookingRequest.MemberID,
				RequestLevel:        thisBooking.BookingRequest.RequestLevel,
				Resource:            thisBooking.BookingRequest.Resource,
				Preference:          thisBooking.BookingRequest.Preference,
				TimeslotID:          val.Booking.TimeslotID,
				StartTime:           val.Booking.StartTime,
				EndTime:             val.Booking.EndTime,
				ApplicationDateTime: time.Now(),
			})
		} else if val.Header == logic.MANAGE {
			if val.Body == logic.TAKE {
				if manageRequest.UserResponses[key] == true { // Change to int later for multiple conflicts
					// Schedule This Booking
					logic.TheBigScheduler(config.Booking{
						MemberID:            thisBooking.BookingRequest.MemberID,
						RequestLevel:        thisBooking.BookingRequest.RequestLevel,
						Resource:            thisBooking.BookingRequest.Resource,
						Preference:          thisBooking.BookingRequest.Preference,
						TimeslotID:          val.Booking.TimeslotID,
						StartTime:           val.Booking.StartTime,
						EndTime:             val.Booking.EndTime,
						ApplicationDateTime: time.Now(),
					})

					db, err := config.Database.Query("DELETE FROM bookings.bookings WHERE bookings.booking_id = $1", val.Conflict.BookingID)
					defer db.Close()
					if err != nil {
						// Probably Do Something
					}

					// TODO Contact About Cancelled Booking
				}
			} else if val.Body == logic.SWITCH {
				if manageRequest.UserResponses[key] == true { // Same as above

					db, err := config.Database.Query("UPDATE bookings.bookings SET bookings.preference TO NULL WHERE bookings.booking_id = $1", val.Conflict.BookingID)
					defer db.Close()
					if err != nil {
						// Probably Do Something
					}

					// Schedule this booking
					logic.TheBigScheduler(config.Booking{
						MemberID:            thisBooking.BookingRequest.MemberID,
						RequestLevel:        thisBooking.BookingRequest.RequestLevel,
						Resource:            thisBooking.BookingRequest.Resource,
						Preference:          thisBooking.BookingRequest.Preference,
						TimeslotID:          val.Booking.TimeslotID,
						StartTime:           val.Booking.StartTime,
						EndTime:             val.Booking.EndTime,
						ApplicationDateTime: time.Now(),
					})
				}
			} else if val.Body == logic.OOF {
				if manageRequest.UserResponses[key] == true { // Guess what, the above is here too
					// Schedule without preference
					logic.TheBigScheduler(config.Booking{
						MemberID:     thisBooking.BookingRequest.MemberID,
						RequestLevel: thisBooking.BookingRequest.RequestLevel,
						Resource:     thisBooking.BookingRequest.Resource,
						// Preference:          thisBooking.BookingRequest.Preference,
						TimeslotID:          val.Booking.TimeslotID,
						StartTime:           val.Booking.StartTime,
						EndTime:             val.Booking.EndTime,
						ApplicationDateTime: time.Now(),
					})
				}
			}
		}
	}

	delete(inProgressBookings, manageRequest.ProgressID)

	fmt.Fprint(w, ":)") // TODO: Proper HTTP Return

}
