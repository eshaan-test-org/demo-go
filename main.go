package main

import "fmt"

const (
	productPath string = "/tmp/products.json"
	sellerPath  string = "/tmp/sellers.json"
)

func main() {
	s := Seller{
		Name: "SomeSeller",
		Address: Address{
			City: "Gotham",
		},
	}
	var i GenericInterface
	i = &s

	givenCity := "Gotham"

	switch i := i.(type) {
	case interface{}:
		fmt.Println("What to do")
	case Product:
		fmt.Println(i.DeliversTo(givenCity))
	case Seller:
		fmt.Println(i.DeliversTo(givenCity))
	}
}

func appendData() {
	var s []string
	s = append(s, productPath)
	s = append(s, sellerPath)
}
