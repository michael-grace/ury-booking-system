package config

import (
	"time"
)

// HTTPRequest is the layout all bodies in an api call are layed out
// Authentication TBD
type HTTPRequest struct {
	messageHead string
	payload     interface{} // Possibly make a specific interface
}

type bookingTimeslots struct {
	itemid    int
	starttime time.Time
	endtime   time.Time
}

// BookingRequest is a request from a user about why and when
// they'd like to book a resource
type BookingRequest struct {
	requestlevel int
	resource     int
	preference   int
	memberid     int
	requests     []bookingTimeslots
}

// CancelRequest is for cancelling items based on booking ID
type CancelRequest struct {
	cancelBookingID []int
}

type movingTimeslots struct {
	bookingID    int
	newStartTime time.Time
	newEndTime   time.Time
}

// MoveRequest is for moving bookings
type MoveRequest struct {
	moveRequests []movingTimeslots
}
