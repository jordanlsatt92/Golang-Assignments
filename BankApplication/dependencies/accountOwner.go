package dependencies

type Owner struct {
	ID         string
	Address    string
	EntityType string
}

func (o *Owner) ChangeAddress(newAddress string) {
	o.Address = newAddress
}
