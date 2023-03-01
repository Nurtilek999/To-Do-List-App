package service

import (
	"Pet_1/internal/domain/entity"
	"Pet_1/internal/domain/repository"
	"errors"
	"fmt"
	"reflect"
)

type ProductService struct {
	productRepo repository.IProductRepository
}

type IProductService interface {
	GetProductByCategory(categ entity.Category) ([]entity.Product, error)
	GetAllCategories() ([]entity.Category, error)
	AddNewProduct(product *entity.ProductViewModel) error
	Edit(product *entity.Product) error
	Delete(product *entity.Product) error
}

func NewProductService(productRepo repository.IProductRepository) *ProductService {
	var productService = ProductService{}
	productService.productRepo = productRepo
	return &productService
}
func (s *ProductService) GetAllCategories() ([]entity.Category, error) {
	rows, err := s.productRepo.GetAllCategories()
	if err != nil {
		return nil, fmt.Errorf("Error in GetAllCategories: %s", err.Error())
	}
	var categories []entity.Category
	for rows.Next() {
		var category entity.Category
		if err = rows.Scan(&category.Name); err != nil {
			return nil, fmt.Errorf("Error in GetAllCategories: %s", err.Error())
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *ProductService) GetProductByCategory(categ entity.Category) ([]entity.Product, error) {
	rows, err := s.productRepo.GetProductsByCategory(categ)
	if err != nil {
		return nil, fmt.Errorf("Error in GetProductByCategory: %s", err.Error())
	}

	var products []entity.Product
	for rows.Next() {
		var prod entity.Product
		if err := rows.Scan(&prod.Name); err != nil {
			return nil, fmt.Errorf("Error in GetProductByCategory: %s", err.Error())
		}
		products = append(products, prod)
	}
	return products, nil
}

func (s *ProductService) AddNewProduct(product *entity.ProductViewModel) error {
	rows, err := s.productRepo.GetCategories()
	if err != nil {
		return fmt.Errorf("Error in GetProductByCategory: %s", err.Error())
	}
	exists := false
	for rows.Next() {
		var categ entity.Category
		if err := rows.Scan(&categ.ID, &categ.Name); err != nil {
			return err
		}
		if categ.ID == product.CategoryID {
			exists = true
			break
		}
	}
	if exists == false {
		err = errors.New("The category does not exist. Create category first")
		return err
	}

	err = s.productRepo.InsertProduct(product)
	return err
}

func (s *ProductService) Edit(product *entity.Product) error {
	rows, err := s.productRepo.GetProduct(product.ID)
	if err != nil {
		return fmt.Errorf("Error in Edit: %s", err.Error())
	}
	var prod entity.Product
	for rows.Next() {
		if err := rows.Scan(&prod.ID, &prod.Name, &prod.Price, &prod.CategoryID); err != nil {
			return fmt.Errorf("Error in Edit: %s", err.Error())
		}
	}
	if product.Name == "" {
		product.Name = prod.Name
	}
	if reflect.ValueOf(product.Price).IsZero() {
		product.Price = prod.Price
	}
	if reflect.ValueOf(product.CategoryID).IsZero() {
		product.CategoryID = prod.CategoryID
	}

	err = s.productRepo.Update(*product)
	if err != nil {
		return fmt.Errorf("Error in Edit: %s", err.Error())
	}

	return nil
}

func (s *ProductService) Delete(product *entity.Product) error {
	err := s.productRepo.Delete(*product)
	if err != nil {
		return fmt.Errorf("Error in Delete: %s", err.Error())
	}
	return nil
}
