package entity

type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

type LoginViewModel struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required,gte=8"`
}
