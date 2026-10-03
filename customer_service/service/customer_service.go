package service

import (
	"customer_service/dto/request"
	"customer_service/dto/response"
)

type CustomerService interface {
	CreateCustomer(customer request.CreateCustomerRequest) (response.CreateCustomerResponse, error)
	FindAllCustomers() ([]response.CreateCustomerResponse, error)
	UpdateCustomer(customerId int64, customer request.UpdateCustomerRequest) (response.CreateCustomerResponse, error)
	DeleteCustomer(customerId int64) (int64, error)
	FindOneCustomer(customerId int64) (response.CreateCustomerResponse, error)
	IsCustomerExist(customerId int64) (bool, error)
}
