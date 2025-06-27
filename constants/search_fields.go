package constants

var (
	EmployeeAllowedFields = map[string]bool{
		"employee_id": true,
		"position":    true,
		"is_active":   true,
	}
	PackageAllowedFields = map[string]bool{
		"package_id":     true,
		"package_type":   true,
		"region_id":      true,
		"package_status": true,
		"registered_at":  true,
	}
	VehicleAllowedFields = map[string]bool{
		"internal_id":        true,
		"vehicle_id":         true,
		"current_load":       true,
		"max_load":           true,
		"led_status":         true,
		"needs_confirmation": true,
		"coord_x":            true,
		"coord_y":            true,
		"AI_coord_x":         true,
		"AI_coord_y":         true,
	}
	RegionAllowedFields = map[string]bool{
		"region_id":        true,
		"region_name":      true,
		"coord_x":          true,
		"coord_y":          true,
		"max_capacity":     true,
		"current_capacity": true,
		"is_full":          true,
		"saturated_at":     true,
	}
	TripLogAllowedFields = map[string]bool{
		"trip_id":     true,
		"vehicle_id":  true,
		"status":      true,
		"destination": true,
		"start_time":  true,
		"end_time":    true,
	}
	DeliveryLogAllowedFields = map[string]bool{
		"trip_id":               true,
		"package_id":            true,
		"region_id":             true,
		"load_order":            true,
		"registered_at":         true,
		"first_transport_time":  true,
		"input_time":            true,
		"second_transport_time": true,
		"completed_at":          true,
	}
)
