package structs

type TotalDataSuccRate struct {
	TotalOrder     float32 `json:"total_order"`
	TotalLeadLog   float32 `json:"total_lead_log"`
	TotalTargetLog float32 `json:"total_target_log"`
}

type TopUserBooking struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	TotalBooking int    `json:"total_booking"`
	//TotalActivity int `json:"total_activity"`
}
