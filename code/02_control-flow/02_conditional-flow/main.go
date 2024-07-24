package main

import "fmt"

func main() {
	for i := 0; i < 13; i++ {
		switch {
		case i == 0:
			fmt.Print("Switch is a conditional ")
		case i == 1:
			fmt.Print("that executes according ")
		case i == 2:
			fmt.Println("to the specified cases.")
		} //   ↓ "i" is the switch statement
		switch i {
		case 3:
			fmt.Print("The switch statement is true by default. ")
		case 4:
			fmt.Print("But it can be anything, ")
		case 5:
			fmt.Println("and it evaluates if switch statement == case statement.")
		}
		switch j := i + 10; j {
		case 16:
			fmt.Print("We can also use an initialization expression before the switch statement. ")
		case 17:
			fmt.Println("initialization expression before the switch statement.")
		case 18:
			fmt.Print("Finally, we have the type switch, ")
			type tipo interface{}
			var x tipo = 10
			switch x.(type) {
			case string:
				fmt.Println("telling me that `i` has a value of type string.") // No...
			case float64:
				fmt.Println("telling me that `i` has a value of type float64.") // No...
			case int:
				fmt.Println("telling me that `i` has a value of type int.") // This runs!
			}
		}
		if i == 9 {
			fmt.Print("\nWe also have if statements, ")
		} else if i == 10 {
			fmt.Print("that follow the pattern of if, else if, else. ")
		} else {
			if i == 11 {
				fmt.Println("You can nest multiple ones inside each other, of course.")
			}
		}
		if k := i + 10; k%2 == 0 && k%11 == 0 {
			fmt.Println("And we can have an initialization expression, and we can chain more than one condition using conditional logical operators (&&, ||, and family).")
		}
	}
}
