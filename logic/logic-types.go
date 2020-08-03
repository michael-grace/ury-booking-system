package logic

import "github.com/michael-grace/ury-booking-system/config"

type ManageType struct {
	Header   string
	Body     string
	Booking  config.BookingTimeslots
	Conflict config.Booking
}

const (
	// ACCEPT is for Accepted Bookings
	ACCEPT = "Accepted Booking"

	// MANAGE is for bookings that need modification
	MANAGE = "Booking Requires Modification"

	// REJECT is for Rejected Bookings
	REJECT = "Rejected Booking"

	// TAKE is for taking a resource of someone else
	TAKE = "Take Resource"

	// SWITCH is for switching someone elses resource preference
	SWITCH = "Switch Resource Preference"

	// OOF is for not being able to get the preference you want
	OOF = "Resource Preference Unavailable"
)
