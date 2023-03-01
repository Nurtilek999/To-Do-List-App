package handler

import (
	"Pet_1/internal/domain/entity"
	"Pet_1/internal/domain/service"
	"Pet_1/pkg/response"
	"Pet_1/pkg/validation"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type ProductHandler struct {
	productService service.IProductService
}

func NewProductHandler(productService service.IProductService) *ProductHandler {
	var productHandler = ProductHandler{}
	productHandler.productService = productService
	return &productHandler
}
func (h *ProductHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.productService.GetAllCategories()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if categories == nil {
		categories = []entity.Category{}
	}

	response.ResponseOKWithData(c, categories)
}
func (h *ProductHandler) GetProductsByCategory(c *gin.Context) {
	var categ entity.Category
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(jsonData, &categ)

	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	
	validateCategory := validation.Validate(&categ)
	if validateCategory != nil {
		response.ResponseError(c, validateCategory.Error(), http.StatusUnprocessableEntity)
		return
	}

	products, err := h.productService.GetProductByCategory(categ)

	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if products == nil {
		products = []entity.Product{}
	}

	response.ResponseOKWithData(c, products)
}

func (h *ProductHandler) Edit(c *gin.Context) {
	var product entity.Product
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(jsonData, &product)

	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	validateProduct := validation.Validate(&product)
	if validateProduct != nil {
		response.ResponseError(c, validateProduct.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.productService.Edit(&product)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

}

func (h *ProductHandler) Delete(c *gin.Context) {
	var product entity.Product
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(jsonData, &product)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	validateProduct := validation.Validate(&product)
	if validateProduct != nil {
		response.ResponseError(c, validateProduct.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.productService.Delete(&product)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

}
func (h *ProductHandler) AddNewProduct(c *gin.Context) {
	var product entity.ProductViewModel
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(jsonData, &product)

	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	validateProduct := validation.Validate(&product)
	if validateProduct != nil {
		response.ResponseError(c, validateProduct.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.productService.AddNewProduct(&product)

	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

}
