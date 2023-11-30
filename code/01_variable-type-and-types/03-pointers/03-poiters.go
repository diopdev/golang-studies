package main

import "fmt"

func main() {
	variable := 10
	fmt.Printf("variable = %v\n", variable)
	// "&"mostra o endereço da variável

	fmt.Printf("&variable = %v\n", &variable)

	adress := &variable
	fmt.Printf("adress = %v\n", adress)
	// "*" faz o de-reference, mostrando o conteúdo da localização pra onde o endereço aponta.
	*adress++
	fmt.Printf("variable + 1 = %v\n", *adress)
	//  É possivel usar os dois ao mesmo tempo.
	//                              ↓↓
	fmt.Printf("*&variable = %v\n", *&variable)

}
