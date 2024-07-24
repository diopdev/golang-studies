package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Thare's a for loop") // Print this message 10 times
	}
	// for loop with a single condition
	i := 0
	for i < 3 {
		fmt.Println("For loop iteration. (This is a while loop in go)") // Print this message 10 times
		i++
	}
	for {
		fmt.Println("And this is an infinite loop. Or it would be, if the wasn't `break`.") // Print this message 10 times
		i++
		if i > 5 {
			break
		}
	}
	for {
		fmt.Println("With a continue statement, we can skip the rest of the loop and start the next iteration.")
		i++
		if i < 9 {
			continue
		}
		fmt.Println("This is the last iteration of the loop.")
		return
	}

}
