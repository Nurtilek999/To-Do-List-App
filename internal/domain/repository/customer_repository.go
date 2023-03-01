package repository

import (
	"Pet_1/internal/domain/entity"
	"database/sql"
)

type CustomerRepository struct {
	db *sql.DB
}

type ICustomerRepository interface {
	GetCustomerByLoginPassword(customerLoginVM entity.CustomerLoginViewModel) (*sql.Rows, error)
	UpdateKorzina(korzina entity.Korzina) error
	InsertKorzina(korzina entity.Korzina) error
	GetKorzinaByCustomerID(id int) (*sql.Rows, error)
	Purchase(id int) error
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	var customerRepo = CustomerRepository{}
	customerRepo.db = db
	return &customerRepo
}

func (r *CustomerRepository) GetCustomerByLoginPassword(customerLoginVM entity.CustomerLoginViewModel) (*sql.Rows, error) {
	rows, err := r.db.Query(`select * from customers where phone_number = $1`, customerLoginVM.PhoneNumber)

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *CustomerRepository) UpdateKorzina(korzina entity.Korzina) error {
	_, err := r.db.Exec(`update korzina set count = $1 where customer_id = $2 and product_id = $3`, korzina.Count, korzina.CustomerID, korzina.ProductID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) InsertKorzina(korzina entity.Korzina) error {
	_, err := r.db.Exec(`insert into korzina values($1, $2, $3)`, korzina.CustomerID, korzina.ProductID, korzina.Count)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) GetKorzinaByCustomerID(id int) (*sql.Rows, error) {
	rows, err := r.db.Query(`select * from korzina where customer_id = $1`, id)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *CustomerRepository) Purchase(id int) error {
	_, err := r.db.Exec(`delete from korzina where customer_id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
