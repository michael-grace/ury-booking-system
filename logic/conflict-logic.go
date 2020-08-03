package logic

import (
	"github.com/michael-grace/ury-booking-system/config"
)

func getNumberOfResources(resource int) (int, error) {
	numberOfResourcesQuery := "SELECT COUNT(unique_resources.unique_id) FROM bookings.unique_resources WHERE unique_resources.resource_id = $1"
	rows, err := config.Database.Query(numberOfResourcesQuery, resource)
	defer rows.Close()
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		var num int
		rows.Scan(&num)
		return num, nil
	}
	return 0, nil
}

func overwriteBookings(full bool, requestGeneral config.BookingRequest, requestSpecific config.BookingTimeslots, conflicts []config.Booking) config.ManageType {

	// This currently only works with 1 conflicting booking
	// Not too difficult to expand (which should get done)
	con := conflicts[0]

	if full {

		if requestGeneral.RequestLevel > con.RequestLevel {
			return config.ManageType{
				Header:   MANAGE,
				Body:     TAKE,
				Booking:  requestSpecific,
				Conflict: con,
			}
		}
		return config.ManageType{Header: REJECT, Booking: requestSpecific}

	}

	if requestGeneral.RequestLevel > con.RequestLevel {
		// Ask to Switch
		return config.ManageType{
			Header:   MANAGE,
			Body:     SWITCH,
			Booking:  requestSpecific,
			Conflict: con,
		}
	}

	// Can't Have Preference
	return config.ManageType{
		Header:   MANAGE,
		Body:     OOF,
		Booking:  requestSpecific,
		Conflict: con,
	}

}

// DealWithConflicts holds the logic for determining where priorities are in bookings
func DealWithConflicts(request config.BookingRequest, conflicts [][]config.Booking) ([]config.ManageType, error) {
	var bookingConflictStatus []config.ManageType

	resourceNum, err := getNumberOfResources(request.Resource)

	if err != nil {
		return nil, err
	}

	for key, val := range request.Requests {

		if len(conflicts[key]) == 0 {
			bookingConflictStatus = append(bookingConflictStatus, config.ManageType{Header: ACCEPT, Booking: val})
			continue
		} else if len(conflicts[key]) == resourceNum {
			bookingConflictStatus = append(bookingConflictStatus, overwriteBookings(true, request, val, conflicts[key]))
		} else {
			preferenceInConflicts := false
			for _, con := range conflicts[key] {
				if request.Preference == con.Preference {
					preferenceInConflicts = true
					break
				}
			}
			if !preferenceInConflicts {
				bookingConflictStatus = append(bookingConflictStatus, config.ManageType{Header: ACCEPT, Booking: val})
			} else {
				bookingConflictStatus = append(bookingConflictStatus, overwriteBookings(false, request, val, conflicts[key]))
			}
		}

	}

	return bookingConflictStatus, nil

}
