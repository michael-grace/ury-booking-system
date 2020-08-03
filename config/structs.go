package config

import (
	"time"
)

// HTTPRequest is the layout all bodies in an api call are layed out
// Authentication TBD
type HTTPRequest struct {
	MessageHead string
	Payload     interface{} // Possibly make a specific interface
}

// Booking is an actual booking in the system
type Booking struct {
	BookingID           int
	MemberID            int
	RequestLevel        int
	Resource            int
	Preference          int
	GivenResource       int
	TimeslotID          int
	StartTime           time.Time
	EndTime             time.Time
	PublicID            int
	ApplicationDateTime time.Time
}

type BookingTimeslots struct {
	TimeslotID int
	StartTime  time.Time
	EndTime    time.Time
}

// BookingRequest is a request from a user about why and when
// they'd like to book a resource
type BookingRequest struct {
	RequestLevel int
	Resource     int
	Preference   int
	MemberID     int
	Requests     []BookingTimeslots
}

// CancelRequest is for cancelling items based on booking ID
type CancelRequest struct {
	CancelBookingID []int
}

type movingTimeslots struct {
	BookingID    int
	newStartTime time.Time
	newEndTime   time.Time
}

// MoveRequest is for moving bookings
type MoveRequest struct {
	MoveRequests []movingTimeslots
}

type ManageType struct {
	Header   string
	Body     string
	Booking  BookingTimeslots
	Conflict Booking
}

// InProgressBooking is for a booking that is waiting for user confirmation
type InProgressBooking struct {
	ProgressID     int
	BookingRequest BookingRequest
	ManageType     []ManageType
}

type ManageRequest struct {
	ProgressID    int
	UserResponses []interface{}
}
