package config

import (
	"time"
)

// Booking is an actual booking in the system
type Booking struct {
	BookingID           int       `json:"bookingID"`
	MemberID            int       `json:"memberID"`
	RequestLevel        int       `json:"requestLevel"`
	Resource            int       `json:"resource"`
	Preference          int       `json:"preference"`
	GivenResource       int       `json:"givenResource"`
	TimeslotID          int       `json:"timeslotID"`
	StartTime           time.Time `json:"startTime"`
	EndTime             time.Time `json:"endTime"`
	PublicID            int       `json:"publicID"`
	ApplicationDateTime time.Time `json:"applicationDateTime"`
}

type BookingTimeslots struct {
	TimeslotID int       `json:"timeslotID"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
}

// BookingRequest is a request from a user about why and when
// they'd like to book a resource
type BookingRequest struct {
	RequestLevel int                `json:"requestLevel"`
	Resource     int                `json:"resource"`
	Preference   int                `json:"preference"`
	MemberID     int                `json:"memberID"`
	Requests     []BookingTimeslots `json:"requests"`
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
	Header   string           `json:"header"`
	Body     string           `json:"body"`
	Booking  BookingTimeslots `json:"booking"`
	Conflict *Booking         `json:"conflict"`
}

// InProgressBooking is for a booking that is waiting for user confirmation
type InProgressBooking struct {
	ProgressID     int            `json:"progressID"`
	BookingRequest BookingRequest `json:"bookingRequest"`
	ManageType     []ManageType   `json:"manageType"`
}

type ManageRequest struct {
	ProgressID    int           `json:"progressID"`
	UserResponses []interface{} `json:"userResponses"`
}

type BookingGetRequest struct {
	DateStart time.Time `json:"dateStart"`
	Resource  string    `json:"resource"`
	Bookings  []Booking `json:"bookings"`
}
