package logic

import (
	"github.com/michael-grace/ury-booking-system/config"
)

// TheBigScheduler schedules things in the DB given a booking b
func TheBigScheduler(b config.Booking) error {

	scheduleQuery := `INSERT INTO bookings.bookings 
	(member_id, request_level, resource, preference, given_resource, timeslot_id, start_time, end_time, public_id, application_datetime)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`

	_, err := config.Database.Query(scheduleQuery,
		b.MemberID,
		b.RequestLevel,
		b.Resource,
		b.Preference,
		b.GivenResource,
		b.TimeslotID,
		b.StartTime,
		b.EndTime,
		b.PublicID,
		b.ApplicationDateTime)

	return err

	// TODO: Allocations

}
