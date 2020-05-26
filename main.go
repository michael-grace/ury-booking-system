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
		Opens and sets up log file
	*/
	logFile, err := os.OpenFile(config.Config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	/*
		Routes HTTP Calls
	*/

	http.HandleFunc("/get", api.GetHandler)
	http.HandleFunc("/add", api.AddHandler)
	http.HandleFunc("/cancel", api.CancelHandler)
	http.HandleFunc("/move", api.MoveHandler)

	/*
		Starts HTTP Server
	*/
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
	if err != nil {
		panic(err)
	}

}
