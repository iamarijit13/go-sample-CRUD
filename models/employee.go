package models

type Employee struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	CTC   float32 `json:"ctc"`
}
