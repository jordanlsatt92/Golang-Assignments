package dependencies

type Owner struct {
	ID         string
	Address    string
	EntityType string
}

// Changes the Owner's Address
func (o *Owner) ChangeAddress(newAddress string) {
	o.Address = newAddress
}
