package logic

import (
	"github.com/michael-grace/ury-booking-system/config"
)

func getNumberOfResources(resource int) (int, error) {
	numberOfResourcesQuery := "SELECT booking_resources.stock FROM bookings.booking_resources WHERE booking_resources.resource_id = $1 AND unique = false;"
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

func getUniqueNumberOfResources(resource int) (int, error) {
	numberOfResourcesQuery := "" // TODO
	rows, err := config.Database.Query(numberOfResourcesQuery, resource)
	defer rows.Close()
	if err != nil {
		return 0, err
	}
	rows.Next()
	var num int
	rows.Scan(&num)
	return num, nil

}

// DealWithConflicts holds the logic for determining where priorities are in bookings
func DealWithConflicts(request config.BookingRequest, conflicts [][]config.Booking) ([]string, error) {
	var bookingConflictStatus []string

	for key, val := range request.Requests {
		if len(conflicts[key]) == 0 {
			bookingConflictStatus = append(bookingConflictStatus, ACCEPT)
			continue
		}

		resourceNum, err := getNumberOfResources(request.Resource)

		if err != nil {
			return nil, err
		}

		if resourceNum > 0 {
			if len(conflicts[key]) < resourceNum {
				bookingConflictStatus = append(bookingConflictStatus, ACCEPT)
				continue
			} else {
				bookingConflictStatus = append(bookingConflictStatus, REJECT)
				continue
			}
		}

		uniqueNum, err := getUniqueNumberOfResources(request.Resource)

		if err != nil {
			return nil, err
		}

		if len(conflicts[key]) >= uniqueNum {
			bookingConflictStatus = append(bookingConflictStatus, REJECT)
			continue
		}

		if request.Preference == 0 {
			bookingConflictStatus = append(bookingConflictStatus, ACCEPT)
			continue
		}

		// TODO

		/*
			Thoughts:
			1) It doesn't matter if someone doesn't have a preference but is a higher priority
			2) People can have higher priorities on non-unique resources
			3) None of this is accounted for
		*/

	}

	return bookingConflictStatus, nil

}
