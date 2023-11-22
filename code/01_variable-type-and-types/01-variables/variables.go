package main

import "fmt"

// Para declaramos variaveis de packege leval scop, utilizamos a keyword "var" ou "const".

var pizza = "Calabresa"

// const sandwich = "Bauru"

func main() {

	//  O gopher operador cria novas variáveis. Ele pode ser usado somente dentro de um scopo local.
	name := "max santos"
	fmt.Println(name)

	// Podemos atualizar o valor de uma variável ja existente usando o operador "="

	name = "Cassy"
	fmt.Println(name)

	fmt.Println(pizza)
}

func food() {
	fmt.Println(pizza)
}
