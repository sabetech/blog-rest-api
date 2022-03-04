package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nerdcademy/restapi/controller/vehiclecontroller"
)

var router *mux.Router

func initHandlers() {
	//I will refactor this bit

	router.HandleFunc("/api/vehicle_categories",
		vehiclecontroller.SetRequiredHeaders(vehiclecontroller.GetAllVehicleCategories)).Methods("GET")
	router.HandleFunc("/api/vehicle_category/{id}/relevant_taxes",
		vehiclecontroller.SetRequiredHeaders(vehiclecontroller.GetVehicleRelevantTaxes)).Methods("GET")
	router.HandleFunc("/api/vehicle_tax",
		vehiclecontroller.SetRequiredHeaders(vehiclecontroller.GetVehicleTotalDuty)).Methods("GET")

}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
