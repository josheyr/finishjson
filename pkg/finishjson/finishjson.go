package finishjson

import (
	"strings"
	"unicode"
)

func pop(stack []rune) []rune {
	if len(stack) > 0 {
		return stack[:len(stack)-1]
	}
	return stack
}

func reverse(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

func FinishJSON(unfinished string) string {
	if len(unfinished) == 0 {
		return "{}"
	}

	// Initialize necessary variables
	var expectedStack []rune
	inString, escaping, expectingValue := false, false, false

	var sb strings.Builder
	sb.WriteString(unfinished)

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
				expectedStack = append(expectedStack, '}')
				expectingValue = false
			case '[':
				expectedStack = append(expectedStack, ']')
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
				expectedStack = append(expectedStack, '"')
				expectingValue = false
			default:
				if len(expectedStack) > 0 && expectedStack[len(expectedStack)-1] == char {
					expectedStack = pop(expectedStack)
				}
				expectingValue = false
			}
		}
	}

	if escaping {
		sb.WriteString("\\")
	}

	// If we are still expecting a value, append "null" to the result
	if expectingValue {
		sb.WriteString("null")
	}

	// Complete the unmatched characters
	if len(expectedStack) > 0 {
		sb.WriteString(string(reverse(expectedStack)))
	}

	return sb.String()
}
