package city

type City struct{
	address string 
}

func New(address string) City {
	return City{
		address: address
	}
}