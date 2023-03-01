package entity

type Customer struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type CustomerLoginViewModel struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,gte=5"`
	Password    string `json:"password" validate:"required,gte=5"`
}

type Korzina struct {
	CustomerID int `json:"customerID"`
	ProductID  int `json:"productID"`
	Count      int `json:"count"`
}

type KorzinaProducts struct {
	List []Korzina `json:"list" validate:"required"`
}
