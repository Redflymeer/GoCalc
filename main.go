package main // Main structure of project

import (
	"fmt"
	"math"
) // Print libary

var first int
var second int
var third int
var more int
var none int


func main() {
	var action string

	fmt.Print("First number (if you don't use, use 0. for / use 1):")
	fmt.Scanln(&first)
	fmt.Print("Second number (if you don't use, use 0. for / use 1):")
	fmt.Scanln(&second)
	fmt.Print("Third number (if you don't use, use 0. for / use 1):")
	fmt.Scanln(&third)
	fmt.Print("Four number (if you don't use, use 0. for / use 1):")
	fmt.Scanln(&more)

	fmt.Print("What action?:")
	fmt.Scanln(&action)

	if first == 0 {
		first = none
	}
	if second == 0 {
		second = none
	}
	if third == 0 {
		third = none
	}
	if more == 0 {
		more = none
	}


	switch action {
	case "+":
		result := first + second + third + more
		fmt.Println("Result:", result)
	case "-":
		result := first - second - third - more
		fmt.Println("Result:", result)
	case "*":
		result := first * second * third * more
		fmt.Println("Result:", result)
	case "/":
   		result := first
    		if second != 0 {
        		result /= second
    		} else {
        		fmt.Println("ERROR: Division by zero (second number)!")
        		return
    		}
    		if third != 0 {
        		result /= third
    		} else {
        		fmt.Println("ERROR: Division by zero (third number)!")
        		return
    		}
    		if more != 0 {
        		result /= more
    		} else {
        		fmt.Println("ERROR: Division by zero (fourth number)!")
        		return
    		}
    		fmt.Println("Result:", result)
	case "**":
    		result := float64(first)
   	 	if second != 0 {
        		result = math.Pow(result, float64(second))
    		}
    		if third != 0 {
        		result = math.Pow(result, float64(third))
    		}
    		if more != 0 {
        		result = math.Pow(result, float64(more))
    		}
    		fmt.Println("Result:", result)
	default:
		fmt.Println("ERROR: Don't actions, use +, -, *, / or **")
	}
}
