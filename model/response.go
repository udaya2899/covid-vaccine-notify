package model

type CowinResponse struct {
	Centers []Centers `json:"centers"`
}
type Sessions struct {
	SessionID         string   `json:"session_id"`
	Date              string   `json:"date"`
	AvailableCapacity int      `json:"available_capacity"`
	MinAgeLimit       int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}
type VaccineFees struct {
	Vaccine string `json:"vaccine"`
	Fee     string `json:"fee"`
}
type Centers struct {
	CenterID     int           `json:"center_id"`
	Name         string        `json:"name"`
	StateName    string        `json:"state_name"`
	DistrictName string        `json:"district_name"`
	BlockName    string        `json:"block_name"`
	Pincode      int           `json:"pincode"`
	Lat          int           `json:"lat"`
	Long         int           `json:"long"`
	From         string        `json:"from"`
	To           string        `json:"to"`
	FeeType      string        `json:"fee_type"`
	Sessions     []Sessions    `json:"sessions"`
	VaccineFees  []VaccineFees `json:"vaccine_fees,omitempty"`
}
