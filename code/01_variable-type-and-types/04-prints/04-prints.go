package main

import "fmt"

func main() {
	// Imprime um testo em stdout
	fmt.Print("fmt.Print joga em os. Stdout. ")

	// Imprime um texto com /n implicito no final. ("/n" = adiciona uma nova linha no final)

	fmt.Println("fmt.Println faz a mesma coisa, mas adiciona uma linha nova.")

	// Copia um texto para uma vaiável do tipo string
	variable := fmt.Sprintf("fmt.Sprintf")

	//  Imprime um texto em stdout substituindo os marcadores % pelos respectivos valores
	fmt.Printf("%v, assim como fmt.Printf, usa format printing. "+
		"Ou seja, toma formatting verbs que podem representar números "+
		"decimais (%d), binários (%b), mostrar tipos (%T), e etc.\n",
		variable, 50, 50, variable)

	// https://golang.org/pkg/fmt/

}
