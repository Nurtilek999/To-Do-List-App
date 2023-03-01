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

type EmployeeHandler struct {
	employeeService service.IEmployeeService
}

func NewEmployeeHandler(employeeService service.IEmployeeService) *EmployeeHandler {
	var employeeHandler = EmployeeHandler{}
	employeeHandler.employeeService = employeeService
	return &employeeHandler
}

func (h *EmployeeHandler) Login(c *gin.Context) {
	var loginVM entity.LoginViewModel

	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(jsonData, &loginVM)

	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	validateEmployee := validation.Validate(&loginVM)
	if validateEmployee != nil {
		response.ResponseError(c, validateEmployee.Error(), http.StatusUnprocessableEntity)
		return
	}

	loginEmployee, err := h.employeeService.GetEmployeeByLoginPassword(loginVM)

	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if loginEmployee == nil {
		loginEmployee = &entity.Employee{}
	}

	response.ResponseOKWithData(c, loginEmployee)
}
