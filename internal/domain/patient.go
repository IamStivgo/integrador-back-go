package domain

type Patient struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DNI 	 string `json:"dni"`
	Address   string `json:"address"`
	Discharge_date string `json:"discharge_date"`
}