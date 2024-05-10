package finishjson

import (
	"strings"
	"unicode"
)

func FinishJSON(unfinished string) string {
	if len(unfinished) == 0 {
		return "{}"
	}

	// Dedicated stack functions for clarity
	push := func(stack []rune, value rune) []rune {
		return append(stack, value)
	}

	pop := func(stack []rune) []rune {
		if len(stack) > 0 {
			return stack[:len(stack)-1]
		}
		return stack
	}

	reverse := func(runes []rune) []rune {
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return runes
	}

	// Initialize necessary variables
	var expectedStack []rune
	inString, escaping, expectingValue := false, false, false

	// Iterate over the input string
	for _, char := range unfinished {
		if inString {
			if escaping {
				escaping = false
			} else if char == '\\' {
				escaping = true
			} else if char == '"' {
				inString = false
				expectedStack = pop(expectedStack)
			}
		} else {
			if unicode.IsSpace(char) {
				continue
			}

			switch char {
			case '{':
				expectedStack = push(expectedStack, '}')
				expectingValue = false
			case '[':
				expectedStack = push(expectedStack, ']')
				expectingValue = true
			case 't':
				expectedStack = append(expectedStack, 'e', 'u', 'r')
				expectingValue = false
			case 'f':
				expectedStack = append(expectedStack, 'e', 's', 'l', 'a')
				expectingValue = false
			case 'n':
				expectedStack = append(expectedStack, 'l', 'l', 'u')
				expectingValue = false
			case ':', ',':
				expectingValue = true
			case '"':
				inString = true
				expectedStack = push(expectedStack, '"')
				expectingValue = false
			default:
				if len(expectedStack) > 0 && expectedStack[len(expectedStack)-1] == char {
					expectedStack = pop(expectedStack)
				}
				expectingValue = false
			}
		}
	}

	unfinished = strings.TrimSpace(unfinished)

	// If we are still expecting a value, append "null" to the result
	if expectingValue {
		unfinished += " null"
	}

	// Complete the unmatched characters
	if len(expectedStack) > 0 {
		unfinished += string(reverse(expectedStack))
	}

	return unfinished
}
