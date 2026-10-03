package repo

import "customer_service/domain"

type CustomerRepository interface {
	CreateCustomer(customer domain.Customer) (domain.Customer, error)
	FindAllCustomers() ([]domain.Customer, error)
	UpdateCustomer(customerId int64, customer domain.Customer) (domain.Customer, error)
	DeleteCustomer(customerId int64) (int64, error)
	FindOneCustomer(customerId int64) (domain.Customer, error)
	IsCustomerExist(customerId int64) (bool, error)
}
