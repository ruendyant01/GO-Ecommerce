package controller

import (
	"customer_service/dto/request"
	"customer_service/dto/response"
	"customer_service/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	service service.CustomerService
}

func NewCustomerController(service service.CustomerService) *CustomerController {
	return &CustomerController{service: service}
}

func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	resp := request.CreateCustomerRequest{}
	err := ctx.ShouldBind(&resp)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse{Code: 400, Message: err.Error()})
		return
	}

	customer, err := c.service.CreateCustomer(resp)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{Code: 500, Message: err.Error()})
		return
	}

	finalResp := response.Response{
		Code:    200,
		Message: "Success Created",
		Data:    customer,
	}

	ctx.JSON(200, finalResp)
}

func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("customerId")
	customerId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse{Code: 400, Message: err.Error()})
		return
	}

	_, err = c.service.FindOneCustomer(int64(customerId))
	if err != nil {
		ctx.JSON(404, response.ErrorResponse{Code: 404, Message: err.Error()})
		return
	}

	_, err = c.service.DeleteCustomer(int64(customerId))
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{Code: 500, Message: err.Error()})
		return
	}

	resp := response.Response{
		Code:    200,
		Message: "Success Deleted",
	}

	ctx.JSON(200, resp)
}

func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
	requestCustomer := request.UpdateCustomerRequest{}
	err := ctx.ShouldBind(&requestCustomer)
	id := ctx.Param("customerId")
	strId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse{Code: 400, Message: err.Error()})
		return
	}
	customer, err := c.service.UpdateCustomer(int64(strId), requestCustomer)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{Code: 500, Message: err.Error()})
		return
	}
	finalResp := response.Response{
		Code:    200,
		Message: "Success Updated",
		Data:    customer,
	}
	ctx.JSON(200, finalResp)
}

func (c *CustomerController) FindAllCustomer(ctx *gin.Context) {
	customers, err := c.service.FindAllCustomers()
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{Code: 500, Message: err.Error()})
		return
	}

	finalResp := response.Response{
		Code:    200,
		Message: "Success",
		Data:    customers,
	}
	ctx.JSON(200, finalResp)
}

func (c *CustomerController) FindOneCustomer(ctx *gin.Context) {
	id := ctx.Param("customerId")
	strId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse{Code: 400, Message: err.Error()})
		return
	}
	customer, err := c.service.FindOneCustomer(int64(strId))
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{Code: 500, Message: err.Error()})
		return
	}
	finalResp := response.Response{
		Code:    200,
		Message: "Success",
		Data:    customer,
	}
	ctx.JSON(200, finalResp)
}
