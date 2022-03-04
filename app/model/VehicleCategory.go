package model

// define vehicle category here
type VehicleCategory struct {
	ID                 uint64 `json:"id"`
	VehicleType        string `json:"vehicle_type"`
	VehicleDescription string `json:"vehicle_desc"`
}

func GetAllVehicleCategories() ([]VehicleCategory, error) {
	var vehicle_categories []VehicleCategory

	query := `select id, vehicle_type, vehicle_desc from vehicle_categories;`

	rows, err := db.Query(query)
	if err != nil {
		return vehicle_categories, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var vehicle_type, vehicle_desc string

		err := rows.Scan(&id, &vehicle_type, &vehicle_desc)
		if err != nil {
			return vehicle_categories, err
		}

		vehicle_category := VehicleCategory{
			ID:                 id,
			VehicleType:        vehicle_type,
			VehicleDescription: vehicle_desc,
		}

		vehicle_categories = append(vehicle_categories, vehicle_category)
	}

	return vehicle_categories, nil
}
