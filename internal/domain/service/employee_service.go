package service

import (
	"Pet_1/internal/domain/entity"
	"Pet_1/internal/domain/repository"
	"errors"
	"fmt"
)

type EmployeeService struct {
	employeeRepo repository.IEmployeeRepository
}

type IEmployeeService interface {
	GetEmployeeByLoginPassword(loginVM entity.LoginViewModel) (*entity.Employee, error)
}

func NewEmployeeService(employeeRepo repository.IEmployeeRepository) *EmployeeService {
	var employeeService = EmployeeService{}
	employeeService.employeeRepo = employeeRepo
	return &employeeService
}

func (s *EmployeeService) GetEmployeeByLoginPassword(loginVM entity.LoginViewModel) (*entity.Employee, error) {
	var check bool
	rows, err := s.employeeRepo.GetEmployeeByLoginPassword(loginVM)
	if err != nil {
		return nil, fmt.Errorf("Error in GetEmployeeByLoginPassword: %s", err.Error())
	}
	var emp entity.Employee
	for rows.Next() {
		if err = rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Login, &emp.Password); err != nil {
			return nil, fmt.Errorf("Error in GetEmployeeByLoginPassword: %s", err.Error())
		}
	}

	// Verify Password
	check = emp.Password == loginVM.Password
	if check == false {
		return nil, errors.New("Incorrect Password")
	}

	return &emp, nil
}
