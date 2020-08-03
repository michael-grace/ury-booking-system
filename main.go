/*
	URY Booking System
	@author: Michael Grace
	@date: 2020
*/

package main

import (
	"fmt"
	"github.com/michael-grace/ury-booking-system/api"
	"github.com/michael-grace/ury-booking-system/config"
	"log"
	"net/http"
	"os"
)

func main() {

	/*
		Sets Config
	*/

	config.ConfigurationSetup()
	defer config.Database.Close()

	/*
		Opens and sets up log file
	*/
	logFile, err := os.OpenFile(config.Config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	/*
		In Progress Bookings
	*/

	var inProgressBookings map[int]config.InProgressBooking

	/*
		Routes HTTP Calls
	*/

	http.HandleFunc("/get", api.GetHandler)
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) { api.AddHandler(w, r, inProgressBookings) })
	http.HandleFunc("/cancel", api.CancelHandler)
	http.HandleFunc("/move", api.MoveHandler)
	http.HandleFunc("/manage", func(w http.ResponseWriter, r *http.Request) { api.ManageHandler(w, r, inProgressBookings) })

	http.HandleFunc("/info/resources", api.ResourceHandler)
	http.HandleFunc("/info/requestlevels", api.PrioritiesHandler)

	/*
		Starts HTTP Server
	*/
	log.Printf("Listening on :%d\n", config.Config.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
	if err != nil {
		panic(err)
	}

}
