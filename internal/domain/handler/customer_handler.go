package handler

import (
	"Pet_1/internal/domain/entity"
	"Pet_1/internal/domain/service"
	"Pet_1/pkg/jwttoken"
	"Pet_1/pkg/response"
	"Pet_1/pkg/validation"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type CustomerHandler struct {
	customerService service.ICustomerService
}

func NewCustomerHandler(customerService service.ICustomerService) *CustomerHandler {
	var customerHandler = CustomerHandler{}
	customerHandler.customerService = customerService
	return &customerHandler
}

func (h *CustomerHandler) Purchase(c *gin.Context) {
	var id int
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(jsonData, &id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.customerService.Purchase(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}
}
func (h *CustomerHandler) EditKorzina(c *gin.Context) {
	var pickedProducts entity.KorzinaProducts
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(jsonData, &pickedProducts)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	validatePickedProducts := validation.Validate(&pickedProducts)
	if validatePickedProducts != nil {
		response.ResponseError(c, validatePickedProducts.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.customerService.EditKorzina(pickedProducts.List)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}
}

func (h *CustomerHandler) CustomerLogin(c *gin.Context) {
	var customerLoginVM entity.CustomerLoginViewModel
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(jsonData, &customerLoginVM)

	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	validateCustomer := validation.Validate(&customerLoginVM)
	if validateCustomer != nil {
		response.ResponseError(c, validateCustomer.Error(), http.StatusUnprocessableEntity)
		return
	}

	token, err := jwttoken.GenerateJWT(customerLoginVM)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	//loginCustomer, err := h.customerService.GetCustomerByLoginPassword(customerLoginVM)
	//
	//if err != nil {
	//	response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
	//	return
	//}
	//
	//if loginCustomer == nil {
	//	loginCustomer = &entity.Customer{}
	//}

	response.ResponseOKWithData(c, token)
}
