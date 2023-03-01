package repository

import (
	"Pet_1/internal/domain/entity"
	"database/sql"
	"errors"
)

type ProductRepository struct {
	db *sql.DB
}

type IProductRepository interface {
	GetAllCategories() (*sql.Rows, error)
	GetProductsByCategory(categ entity.Category) (*sql.Rows, error)
	GetProduct(id int) (*sql.Rows, error)
	Update(product entity.Product) error
	Delete(product entity.Product) error
	GetCategories() (*sql.Rows, error)
	InsertProduct(product *entity.ProductViewModel) error
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	var productRepo = ProductRepository{}
	productRepo.db = db
	return &productRepo
}

func (r *ProductRepository) GetAllCategories() (*sql.Rows, error) {

	rows, err := r.db.Query("select name from categories")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *ProductRepository) GetProductsByCategory(categ entity.Category) (*sql.Rows, error) {
	rows, err := r.db.Query("select p.name from products p left join categories c on p.category_id = c.id where c.name = $1", categ.Name)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
func (r *ProductRepository) GetCategories() (*sql.Rows, error) {
	rows, err := r.db.Query("select * from categories")
	if err != nil {
		return nil, err
	}
	return rows, nil
}
func (r *ProductRepository) InsertProduct(product *entity.ProductViewModel) error {
	_, err := r.db.Exec("insert into products(name, price, category_id) values ($1, $2, $3)", product.Name, product.Price, product.CategoryID) //.Scan(&product.ID)
	if err != nil {
		return err
	}
	return errors.New("The product is added successfully")
}

func (r *ProductRepository) GetProduct(id int) (*sql.Rows, error) {
	rows, err := r.db.Query("select * from products where id = $1", id)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *ProductRepository) Update(product entity.Product) error {
	_, err := r.db.Exec("update products set name = $1, price = $2, category_id = $3 where id = $4", product.Name, product.Price, product.CategoryID, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Delete(product entity.Product) error {
	_, err := r.db.Exec("delete from products where id = $1", product.ID)
	if err != nil {
		return err
	}
	return nil
}
