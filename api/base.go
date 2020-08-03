package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	"io/ioutil"
	"net/http"
)

func baseHTTPRequest(call string, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No Body %v", err)
		return nil, err
	}

	if call == "ADD" {
		var apiRequest config.BookingRequest
		err = json.Unmarshal(reqBody, &apiRequest)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad JSON %v", err)
			return nil, err
		}
		return apiRequest, nil
	} else if call == "CNL" {
		var apiRequest config.CancelRequest
		err = json.Unmarshal(reqBody, &apiRequest)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad JSON %v", err)
			return nil, err
		}
		return apiRequest, nil
	} else if call == "MAN" {
		var apiRequest config.ManageRequest
		err = json.Unmarshal(reqBody, &apiRequest)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad JSON %v", err)
			return nil, err
		}
		return apiRequest, nil
	} else if call == "MOVE" {
		var apiRequest config.MoveRequest
		err = json.Unmarshal(reqBody, &apiRequest)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad JSON %v", err)
			return nil, err
		}
		return apiRequest, nil
	} else {
		fmt.Fprint(w, "No Call")
		return nil, nil
	}

}
