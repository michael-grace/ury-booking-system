package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	"net/http"
)

type Resource struct {
	ResourceID int    `json:"resourceID"`
	Name       string `json:"name"`
}

func InformationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(1)
	rows, err := config.Database.Query("SELECT resources.resource AS ResourceID, resources.name AS Name FROM bookings.resources")
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var resources []Resource
	fmt.Println(2)
	for rows.Next() {
		var resource Resource
		err = rows.Scan(&resource)
		if err != nil {
			panic(err)
		}
		resources = append(resources, resource)
	}

	jsonData, err := json.MarshalIndent(resources, "", "	")

	if err != nil {
		return
	}
	fmt.Println(3)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, jsonData)
}
