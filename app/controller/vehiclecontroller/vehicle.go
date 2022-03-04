package vehiclecontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nerdcademy/restapi/model"
)

func SetRequiredHeaders(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	}
}

// /api/vehicle_categories
func GetAllVehicleCategories(w http.ResponseWriter, r *http.Request) {

	vehicle_categories, err := model.GetAllVehicleCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(vehicle_categories)
	}
}

// /api/vehicle_category/{id}/relevant_taxes
func GetVehicleRelevantTaxes(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query().Get("id") - an alternative
	params := mux.Vars(r)
	vehicleCategoryID := params["id"]
	vehicleTypeId, _ := strconv.Atoi(vehicleCategoryID)
	relevantTaxes, err := model.GetVehicleRelvantTaxes(vehicleTypeId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(relevantTaxes)
	}

}

// /api/vehicle_tax
/*
	Params
		vehicle_category_id: uint
		cif: float
*/
func GetVehicleTotalDuty(w http.ResponseWriter, r *http.Request) {
	vehicle_category_id := r.URL.Query()["vehicle_category_id"]
	cifParam := r.URL.Query()["cif"]

	vehicleCategoryId, _ := strconv.Atoi(vehicle_category_id[0])
	vehicleTaxesDto, _ := model.GetVehicleRelvantTaxes(vehicleCategoryId)
	cif, _ := strconv.ParseFloat(cifParam[0], 32)
	cif32 := float32(cif)
	totalDuty := model.CalculateDuty(vehicleTaxesDto, cif32)

	json.NewEncoder(w).Encode(totalDuty)
}
