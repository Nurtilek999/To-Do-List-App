package service

import (
	"Pet_1/internal/domain/entity"
	"Pet_1/internal/domain/repository"
	"errors"
	"fmt"
)

type CustomerService struct {
	customerRepo repository.ICustomerRepository
}

type ICustomerService interface {
	GetCustomerByLoginPassword(customerLoginVM entity.CustomerLoginViewModel) (*entity.Customer, error)
	EditKorzina(pickedProducts []entity.Korzina) error
	Purchase(id int) error
}

func NewCustomerService(customerRepo repository.ICustomerRepository) *CustomerService {
	var customerService = CustomerService{}
	customerService.customerRepo = customerRepo
	return &customerService
}

func (s *CustomerService) GetCustomerByLoginPassword(customerLoginVM entity.CustomerLoginViewModel) (*entity.Customer, error) {
	rows, err := s.customerRepo.GetCustomerByLoginPassword(customerLoginVM)
	if err != nil {
		return nil, fmt.Errorf("Error in GetCustomerByLoginPassword: %s", err.Error())
	}

	var customer entity.Customer
	for rows.Next() {
		if err := rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.PhoneNumber, &customer.Password); err != nil {
			return nil, err
		}
	}

	if customer.Password != customerLoginVM.Password {
		return nil, errors.New(`Incorrect Phone Number or Password`)
	}
	return &customer, nil

}

func (s *CustomerService) Purchase(id int) error {
	err := s.customerRepo.Purchase(id)
	if err != nil {
		return fmt.Errorf("Error in Purchase: %s", err.Error())
	}
	return nil
}

func (s *CustomerService) EditKorzina(pickedProducts []entity.Korzina) error {
	rows, err := s.customerRepo.GetKorzinaByCustomerID(pickedProducts[0].CustomerID)
	if err != nil {
		return fmt.Errorf("Error in EditKorzina: %s", err.Error())
	}
	var existingProducts []entity.Korzina
	for rows.Next() {
		var korzina entity.Korzina
		if err := rows.Scan(&korzina.CustomerID, &korzina.ProductID, &korzina.Count); err != nil {
			return fmt.Errorf("Error in EditKorzina: %s", err.Error())
		}
		existingProducts = append(existingProducts, korzina)
	}

	i := 0
	j := 0

	for i < len(pickedProducts) {
		for j < len(existingProducts) {
			if pickedProducts[i].ProductID == existingProducts[j].ProductID {
				var currentKorzina entity.Korzina
				currentKorzina.CustomerID = pickedProducts[i].CustomerID
				currentKorzina.ProductID = pickedProducts[i].ProductID
				currentKorzina.Count = pickedProducts[i].Count + existingProducts[j].Count
				err := s.customerRepo.UpdateKorzina(currentKorzina)
				if err != nil {
					return fmt.Errorf("Error in EditKorzina: %s", err.Error())
				}
				i++
				j = 0
				break
			}
			j++
		}
		if j == len(existingProducts) {
			err := s.customerRepo.InsertKorzina(pickedProducts[i])
			if err != nil {
				return fmt.Errorf("Error in EditKorzina: %s", err.Error())
			}
			i++
			j = 0
		}
	}

	return nil
}
