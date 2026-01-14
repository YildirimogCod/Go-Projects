package main

import "fmt"

func main() {
	fmt.Println("Welcome to a simple CLI Calculator!")
	var option int
	result := 0.0
	var number1, number2 float64
	running := true
	for running {
		fmt.Println("\n1.Addition\n2. Substraction\n3. Multiplication\n 4.Division\n 5.Exit")
		fmt.Println("Select an option: ")
		if _, err := fmt.Scanln(&option); err != nil {
			fmt.Println("Invalid Operation")
			continue
		}
		if option <= 0 || option > 5 {
			fmt.Println("Invalid Operation")
			continue
		}
		if option == 5 {
			break
		}
		fmt.Println("Enter the first number")
		if _, err := fmt.Scanln(&number1); err != nil {
			fmt.Println("Invalid Operation")
			continue
		}
		fmt.Println("Enter the second number")
		if _, err := fmt.Scanln(&number2); err != nil {
			fmt.Println("Invalid Operation")
			continue
		}
		if option == 4 && option == 0 {
			fmt.Println("Cannot divide zero!")
			continue
		}

		switch option {
		case 1:
			result = number1 + number2
			fmt.Printf("%g + %g = %g", number1, number2, result)
		case 2:
			result = number1 - number2
			fmt.Printf("%g - %g = %g", number1, number2, result)
		case 3:
			result = number1 * number2
			fmt.Printf("%g * %g = %g", number1, number2, result)
		case 4:
			result = number1 / number2
			fmt.Printf("%g / %g = %g", number1, number2, result)
		}

	}
	fmt.Println("Take care....")
}
