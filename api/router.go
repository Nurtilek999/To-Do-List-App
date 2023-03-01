package api

import (
	"Pet_1/api/middleware"
	"Pet_1/internal/domain/handler"
	"Pet_1/internal/domain/repository"
	"Pet_1/internal/domain/service"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {

	router := gin.Default()

	//Register User Repo
	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	customerRepo := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)

	employee := router.Group("v1/employee")
	{
		employee.POST("/login", employeeHandler.Login)
	}

	product := router.Group("v1/product")
	{
		product.POST("/products", productHandler.GetProductsByCategory)
		product.GET("/categories", productHandler.GetAllCategories)
		product.PUT("/add_product", productHandler.AddNewProduct)
		product.POST("/update", productHandler.Edit)
		product.POST("/delete", productHandler.Delete)
	}

	customer := router.Group("v1/customer")
	{
		customer.POST("/login", customerHandler.CustomerLogin)
		customer.POST("/editKorzina", middleware.AuthMiddleware(), customerHandler.EditKorzina)
		customer.POST("/purchase", customerHandler.Purchase)
	}

	return router

}
