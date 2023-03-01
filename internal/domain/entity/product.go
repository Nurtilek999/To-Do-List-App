package entity

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	CategoryID int     `json:"category_id"`
}

type ProductViewModel struct {
	//ID         int     `json:"id"`
	Name       string  `json:"name" validate:"required"`
	Price      float32 `json:"price"`
	CategoryID int     `json:"category_id"`
}

type Category struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name"`
}
