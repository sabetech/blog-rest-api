package model

import "fmt"

// define vehicle import tax here

type VehicleImportTaxDto struct {
	VehicleType              string  `json:"type" pg:"vehicle_type"`
	VehicleDescription       string  `json:"desc" pg:"vehicle_desc"`
	VehicleImportDuty        float32 `json:"importDuty" pg:"import_duty"`
	VehicleVat               float32 `json:"vat" pg:"vat"`
	VehicleNHIL              float32 `json:"nhil" pg:"nhil"`
	VechicleGetfundLevy      float32 `json:"getfundLevy" pg:"getfund_levy"`
	VehicleAu_levy           float32 `json:"auLevy" pg:"au_levy"`
	VehicleEcowas            float32 `json:"ecowas" pg:"ecowas"`
	VehicleEximLevy          float32 `json:"eximLevy" pg:"exim_levy"`
	VehicleExamFee           float32 `json:"examFee" pg:"exam_fee"`
	VehicleProcessFee        float32 `json:"processFee" pg:"process_fee"`
	VehicleSpecialImportLevy float32 `json:"specialImportLevy" pg:"special_import_levy"`
	VehicleCategoryID        int64   `json:"vehicleCategory" pg:"vehicle_categories_id"`
}

func GetVehicleRelvantTaxes(vehicleCategoryID int) (*VehicleImportTaxDto, error) {
	//make a query to get the relevant taxes on this vehicle
	// query := fmt.Sprintln("select * from vehicle_taxes where vehicle_categories_id = %v", vehicleCategoryID)
	// query2 := fmt.Sprintln("select * from vehicle_categories where id = &v", vehicleCategoryID)

	query := fmt.Sprintln(`select b.vehicle_type, b.vehicle_desc, a.import_duty, a.vat, a.nhil, a.getfund_levy,
							a.au_levy,
							a.ecowas,
							a.exim_levy,
							a.process_fee,
							a.special_import_levy,
							b.id vehicle_category_id,
							from vehicle_taxes a
							join vehicle_categories b on b.id = a.vehicle_categories_id
							where vehicle_categories_id = ?`)

	var vehicleImportTaxDto VehicleImportTaxDto
	row, err := db.Query(query, vehicleCategoryID)
	if err != nil {
		return nil, err
	}

	row.Scan(&vehicleImportTaxDto)

	return &vehicleImportTaxDto, nil

}

func CalculateDuty(vehicleImportInfo *VehicleImportTaxDto, CIF float32) float32 {
	vat := CIF + CIF*vehicleImportInfo.VehicleImportDuty + (CIF * vehicleImportInfo.VehicleNHIL) +
		(CIF+vehicleImportInfo.VechicleGetfundLevy)*vehicleImportInfo.VehicleVat

	return CIF * (vehicleImportInfo.VehicleImportDuty + vehicleImportInfo.VehicleNHIL +
		vehicleImportInfo.VehicleAu_levy + vehicleImportInfo.VehicleEcowas + vehicleImportInfo.VechicleGetfundLevy +
		vehicleImportInfo.VehicleEximLevy + vehicleImportInfo.VehicleExamFee +
		vehicleImportInfo.VehicleProcessFee + vehicleImportInfo.VehicleSpecialImportLevy +
		vat)
}
