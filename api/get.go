package api

import (
	"encoding/json"
	"github.com/michael-grace/ury-booking-system/config"
	"net/url"
	"time"
	// "io/ioutil"
	"fmt"
	"net/http"
)

// GetHandler deals with returning bookings within a timeframe in the body
func GetHandler(w http.ResponseWriter, r *http.Request) {

	getData, err := url.ParseQuery(r.URL.String()[1:])

	if err != nil {
		fmt.Fprintf(w, "Bad GET Request: %v", err)
		return
	}

	reqResourceQuery, ok := getData["resource"]
	if !ok {
		fmt.Fprint(w, "No Resource Specified")
		return
	}
	if len(reqResourceQuery) != 1 {
		fmt.Fprint(w, "Bad Resource Query")
		return
	}
	reqResource := reqResourceQuery[0]

	reqDateStringQuery, ok := getData["date"]
	if !ok {
		fmt.Fprint(w, "No Date Specified")
		return
	}
	if len(reqDateStringQuery) != 1 {
		fmt.Fprint(w, "Bad Date Query")
		return
	}
	reqDateString := reqDateStringQuery[0]
	reqDate, err := time.Parse("2006-01-02", reqDateString)
	if err != nil {
		fmt.Fprintf(w, "Bad Date Conversion: %v", err)
		return
	}

	rows, err := config.Database.Query(`SELECT * FROM bookings.bookings
	INNER JOIN bookings.resources USING (resource)
	WHERE (bookings.start_time >= $1 AND bookings.end_time <= $2)
	OR (bookings.end_time >= $1 AND bookings.end_time <= $2)
	AND resources.name = $3;`, reqDate, reqDate.Add(24*time.Hour), reqResource)
	defer rows.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Database Query Failed %v \n", err)
		return
	}

	var bookings []config.Booking

	for rows.Next() {
		var booking config.Booking
		err = rows.Scan(
			&booking.BookingID,
			&booking.MemberID,
			&booking.RequestLevel,
			&booking.Resource,
			&booking.Preference,
			&booking.GivenResource,
			&booking.TimeslotID,
			&booking.StartTime,
			&booking.EndTime,
			&booking.PublicID,
			&booking.ApplicationDateTime,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Database Output Flawed")
			return
		}

		bookings = append(bookings, booking)
	}

	toReturn := config.BookingGetRequest{
		DateStart: reqDate,
		Resource:  reqResource,
		Bookings:  bookings,
	}

	jsonData, err := json.MarshalIndent(toReturn, "", "	")

	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))

}
