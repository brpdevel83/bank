package service

import domain "github.com/brpdevel83/bank/domian"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaualtCustomerService struct {
	repo domain.CustomerRepository
}

func GetAllCustomers(s DefaualtCustomerService) ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaualtCustomerService {
	return DefaualtCustomerService{repository}
}
