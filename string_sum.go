package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	array []string
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	array = strings.Split(input, "+")
	if len(array) == 1 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	if len(array) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	return COMPUTE(array[0], array[1])
}
func OPERANDS(string1, string2 string) bool {
	string1Minus := strings.LastIndex(string1, "-")
	if string1Minus > 0 {
		return false
	}
	string2Minus := strings.LastIndex(string2, "-")
	if string2Minus > 0 {
		return false
	}
	Plus1 := strings.LastIndex(string1, "+")
	if Plus1 > 0 {
		return true
	}
	Plus2 := strings.LastIndex(string2, "+")
	return Plus2 <= 0
}
func COMPUTE(string1, string2 string) (string,error) {
	if !OPERANDS(string1, string2) {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	op1, error1 := strconv.Atoi(string1)
	if error1 != nil {
		return "", fmt.Errorf("%w", error1)
	}
	op2, error2 := strconv.Atoi(string2)
	if error2 != nil {
		return "", fmt.Errorf("%w", error2)
	}
	return strconv.Itoa(op1+op2), nil
}