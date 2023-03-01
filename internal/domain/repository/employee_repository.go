package repository

import (
	"Pet_1/internal/domain/entity"
	"database/sql"
)

type EmployeeRepository struct {
	db *sql.DB
}

type IEmployeeRepository interface {
	GetEmployeeByLoginPassword(loginVM entity.LoginViewModel) (*sql.Rows, error)
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	var employeeRepo = EmployeeRepository{}
	employeeRepo.db = db
	return &employeeRepo
}

func (r *EmployeeRepository) GetEmployeeByLoginPassword(loginVM entity.LoginViewModel) (*sql.Rows, error) {
	rows, err := r.db.Query("select id, first_name, last_name, login, password from employees where login = $1", loginVM.Login)

	if err != nil {
		return nil, err
	}
	return rows, nil
}
