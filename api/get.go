package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	"io/ioutil"
	"net/http"
)

// GetHandler deals with returning bookings within a timeframe in the body
func GetHandler(w http.ResponseWriter, r *http.Request) {

	apiRequest, err := baseHTTPRequest(w, r)
	if err != nil {
		return
	}

}
