package main

import "fmt"

func main() {
	variable := 10
	fmt.Printf("variable = %v\n", variable)
	// mostra o endereço da variável

	fmt.Printf("&variable = %v\n", &variable)

	adress := &variable
	fmt.Printf("adress = %v\n", adress)
}
