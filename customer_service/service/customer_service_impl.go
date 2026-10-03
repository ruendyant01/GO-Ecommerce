package service

import (
	"customer_service/domain"
	"customer_service/dto/request"
	"customer_service/dto/response"
	"customer_service/repo"
	"strconv"
)

type CustomerServiceImpl struct {
	repository repo.CustomerRepository
}

func NewCustomerServiceImpl(repo repo.CustomerRepository) *CustomerServiceImpl {
	return &CustomerServiceImpl{
		repository: repo,
	}
}

func (c CustomerServiceImpl) CreateCustomer(customer request.CreateCustomerRequest) (response.CreateCustomerResponse, error) {
	newCustomer := domain.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Address: domain.Address{
			Street:      customer.Address.Street,
			HouseNumber: customer.Address.HouseNumber,
			ZipCode:     customer.Address.ZipCode,
		},
	}
	createCustomer, err := c.repository.CreateCustomer(newCustomer)
	if err != nil {
		return response.CreateCustomerResponse{}, err
	}

	return response.CreateCustomerResponse{
		CustomerId: strconv.FormatInt(createCustomer.Id, 10),
		FirstName:  createCustomer.FirstName,
		LastName:   createCustomer.LastName,
	}, nil
}

func (c CustomerServiceImpl) FindAllCustomers() ([]response.CreateCustomerResponse, error) {
	results, err := c.repository.FindAllCustomers()
	if err != nil {
		return nil, err
	}

	var responses []response.CreateCustomerResponse
	for _, result := range results {
		responses = append(responses, response.CreateCustomerResponse{
			CustomerId: strconv.FormatInt(result.Id, 10),
			FirstName:  result.FirstName,
			LastName:   result.LastName,
		})
	}

	return responses, nil
}

func (c CustomerServiceImpl) UpdateCustomer(customerId int64, customer request.UpdateCustomerRequest) (response.CreateCustomerResponse, error) {
	newUpdateCustomer := domain.Customer{
		Id:        customerId,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Address: domain.Address{
			Street:      customer.Address.Street,
			HouseNumber: customer.Address.HouseNumber,
			ZipCode:     customer.Address.ZipCode,
		},
	}

	updateCustomer, err := c.repository.UpdateCustomer(customerId, newUpdateCustomer)
	if err != nil {
		return response.CreateCustomerResponse{}, err
	}
	return response.CreateCustomerResponse{
		CustomerId: strconv.FormatInt(updateCustomer.Id, 10),
		FirstName:  updateCustomer.FirstName,
		LastName:   updateCustomer.LastName,
	}, nil
}

func (c CustomerServiceImpl) DeleteCustomer(customerId int64) (int64, error) {
	_, err := c.repository.DeleteCustomer(customerId)
	if err != nil {
		return 0, err
	}
	return customerId, nil
}

func (c CustomerServiceImpl) FindOneCustomer(customerId int64) (response.CreateCustomerResponse, error) {
	customer, err := c.repository.FindOneCustomer(customerId)
	if err != nil {
		return response.CreateCustomerResponse{}, err
	}
	return response.CreateCustomerResponse{
		CustomerId: strconv.FormatInt(customer.Id, 10),
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
	}, nil
}

func (c CustomerServiceImpl) IsCustomerExist(customerId int64) (bool, error) {
	return c.repository.IsCustomerExist(customerId)
}
