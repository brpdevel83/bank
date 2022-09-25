type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) finadAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customer := []Customer{
		{1, "babak", "Istanbul", "1983-07-01", 1},
		{1, "Nuğdiş", "Istanbul", "1981-01-07", 1},
	}
	return CustomerRepositoryStub{customer}
}