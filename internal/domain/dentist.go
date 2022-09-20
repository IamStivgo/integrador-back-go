package domain

type Dentist struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Enrollment string `json:"enrollment"`
}
