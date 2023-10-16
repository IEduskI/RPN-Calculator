package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Operation interface {
	Execute(symbol string, op ...int) (int, error)
}

func evaluateRPN(expression string) (int, error) {
	tokens := strings.Fields(expression)
	stack := make([]int, 0)

	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid RPN expression")
			}

			// Pop the last two operands
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Perform the operation based on the operator
			switch token {
			case "+":
				stack = append(stack, op1+op2)
			case "-":
				stack = append(stack, op1-op2)
			case "*":
				stack = append(stack, op1*op2)
			case "/":
				if op2 == 0 {
					return 0, fmt.Errorf("division by zero")
				}

				stack = append(stack, op1/op2)
			default:
				return 0, fmt.Errorf("unknown operator: %s", token)
			}
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid RPN expression")
	}

	return stack[0], nil
}

func main() {
	expression := "2 3 4 * +"
	result, err := evaluateRPN(expression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
