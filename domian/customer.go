package domain

type Customer struct {
	Id          int
	Name        string
	City        string
	DateOfBirth string
	status      int
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
