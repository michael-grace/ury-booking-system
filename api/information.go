package api

import (
	"encoding/json"
	"fmt"
	"github.com/michael-grace/ury-booking-system/config"
	"net/http"
)

// ResourceReturn is sent to the user containing all resources available
type ResourceReturn struct {
	Resources []Resource `json:"resources"`
}

// Resource is an individual type of resource
type Resource struct {
	ResourceID      int              `json:"resourceID"`
	Name            string           `json:"name"`
	UniqueResources []UniqueResource `json:"uniqueResources"`
}

// UniqueResource is an individual resource, which will be part of a type
type UniqueResource struct {
	UniqueID int    `json:"uniqueID"`
	Name     string `json:"name"`
}

// ResourceHandler deals with requests for the resource information
func ResourceHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := config.Database.Query("SELECT * FROM bookings.resources")
	defer rows.Close()
	if err != nil {
		fmt.Fprint(w, err)
	}

	var resources []Resource
	for rows.Next() {
		var resource Resource
		err = rows.Scan(&resource.ResourceID, &resource.Name)
		if err != nil {
			fmt.Fprint(w, err)
		}

		// Individual Resources
		rows2, err := config.Database.Query("SELECT unique_resources.unique_id, unique_resources.name FROM bookings.unique_resources WHERE unique_resource.resource = $1", resource.ResourceID)
		defer rows2.Close()
		if err != nil {
			fmt.Fprint(w, err)
		}

		var uniqueResources []UniqueResource

		for rows2.Next() {
			var unique UniqueResource
			err = rows2.Scan(&unique.UniqueID, &unique.Name)
			if err != nil {
				fmt.Fprint(w, err)
			}
			uniqueResources = append(uniqueResources, unique)
		}
		resource.UniqueResources = uniqueResources
		resources = append(resources, resource)
	}

	jsonData, err := json.MarshalIndent(resources, "", "	")

	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, jsonData)

}
