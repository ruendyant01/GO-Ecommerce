package impl

import (
	"customer_service/domain"
	"customer_service/repo"
	"customer_service/util"
	"errors"
	"log"

	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	Db *gorm.DB
}

func NewCustomerRepositoryImpl(db *gorm.DB) repo.CustomerRepository {
	return &CustomerRepositoryImpl{Db: db}
}

func (c CustomerRepositoryImpl) CreateCustomer(customer domain.Customer) (domain.Customer, error) {
	result := c.Db.Create(&customer)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return customer, nil
}

func (c CustomerRepositoryImpl) FindAllCustomers() ([]domain.Customer, error) {
	var customers []domain.Customer
	result := c.Db.Raw(util.FIND_ALL_DB_QUERY_RAW).Find(&customers)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return customers, nil
}

func (c CustomerRepositoryImpl) UpdateCustomer(customerId int64, customer domain.Customer) (domain.Customer, error) {
	result := c.Db.Where("id = ?", customerId).UpdateColumns(&customer)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return customer, nil
}

func (c CustomerRepositoryImpl) DeleteCustomer(customerId int64) (int64, error) {
	var customer domain.Customer
	result := c.Db.Where("id = ?", customerId).Delete(&customer)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return customerId, nil
}

func (c CustomerRepositoryImpl) FindOneCustomer(customerId int64) (domain.Customer, error) {
	var customer domain.Customer
	result := c.Db.Where("id = ?", customerId).First(&customer)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	if result.RowsAffected == 0 {
		return customer, errors.New("customer not found")
	}
	return customer, nil
}

func (c CustomerRepositoryImpl) IsCustomerExist(customerId int64) (bool, error) {
	customer, err := c.FindOneCustomer(customerId)
	if err != nil {
		log.Fatal(err)
	}

	if customer.Id == 0 {
		return false, nil
	}
	return true, nil
}
