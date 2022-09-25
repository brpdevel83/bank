package service

import domain "github.com/brpdevel83/bank/domian"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, nil)
}
