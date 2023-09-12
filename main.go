package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func isRomanNumeral(s string) bool {
	_, ok := romanNumerals[s]
	return ok
}

func parseRomanNumeral(s string) int {
	return romanNumerals[s]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Simple Calculator")
	fmt.Println("Available operators: +, -, *, /")
	fmt.Println("Enter 'exit' to quit")

	for {
		fmt.Print("Enter an expression: ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting calculator")
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Error: Invalid input format")
			continue
		}

		var num1, num2 int
		var err1, err2 error

		if isRomanNumeral(parts[0]) && isRomanNumeral(parts[2]) {
			num1 = parseRomanNumeral(parts[0])
			num2 = parseRomanNumeral(parts[2])
		} else {
			num1, err1 = strconv.Atoi(parts[0])
			num2, err2 = strconv.Atoi(parts[2])

			if err1 != nil || err2 != nil {
				fmt.Println("Error: Invalid number")
				continue
			}
		}

		operator := parts[1]
		result := 0

		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 != 0 {
				result = num1 / num2
			} else {
				fmt.Println("Error: Division by zero")
				continue
			}
		default:
			fmt.Println("Error: Invalid operator")
			continue
		}

		if isRomanNumeral(parts[0]) && isRomanNumeral(parts[2]) {
			fmt.Printf("Result: %s\n", getRomanNumeral(result))
		} else {
			fmt.Printf("Result: %d\n", result)
		}
	}
}

func getRomanNumeral(n int) string {
	romanNumeral := ""
	for n > 0 {
		for numeral, value := range romanNumerals {
			if value <= n {
				romanNumeral += numeral
				n -= value
				break
			}
		}
	}
	return romanNumeral
}

