package main

import "fmt"

func main() {
	// Should output false false false true
	testClosingBraceStart := "}()()()[]"
	testClosingBraceEnd := "[][]([])}"
	testRandomBraceMiddle := "[[]]}[[]]"
	testValidString := "[][[]](){}{}{}"
	fmt.Println(validParenthesis(testClosingBraceStart))
	fmt.Println(validParenthesis(testClosingBraceEnd))
	fmt.Println(validParenthesis(testRandomBraceMiddle))
	fmt.Println(validParenthesis(testValidString))
}

func validParenthesis(s string) bool {
	stack := []string{}
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		// Everytime we find an opening brace
		// we add a closing brace to the stack
		if char == "{" {
			stack = append(stack, "}")
		} else if char == "(" {
			stack = append(stack, ")")
		} else if char == "[" {
			stack = append(stack, "]")
		} else {
			// If there is no closing brace in the stack
			// And we find a random closing brace
			// Return false
			if len(stack) == 0 {
				return false
			}
			// We check that the closing brace mirrors the last
			// Opening brace
			compare := stack[len(stack)-1]
			// If it does mirror the closing brace we
			// Remove it from the stack
			if char == compare {
				stack = stack[:len(stack)-1]
			} else {
				// If it is the wrong closing brace
				return false
			}
		}
	}
	// If there was somehow a last opening brace
	// The stack will have one extra closing brace
	// We return false
	if len(stack) > 0 {
		return false
	}
	// If it iterates through the string and all checks pass
	return true
}
