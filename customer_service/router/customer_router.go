package router

import (
	"customer_service/controller"

	"github.com/gin-gonic/gin"
)

func CustomerRouter(ctrl *controller.CustomerController) *gin.Engine {
	service := gin.Default()
	router := service.Group("/customer")
	router.GET("", ctrl.FindAllCustomer)
	router.POST("", ctrl.CreateCustomer)
	router.GET("/:customerId", ctrl.FindOneCustomer)
	router.DELETE("/:customerId", ctrl.DeleteCustomer)
	router.PUT("/:customerId", ctrl.UpdateCustomer)

	return service
}
