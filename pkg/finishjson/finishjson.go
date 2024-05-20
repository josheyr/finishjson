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

	var (
		expectedStack  []rune
		inString       bool
		escaping       bool
		expectingValue bool
		lastComma      bool
		sb             strings.Builder
	)

	sb.WriteString(unfinished)

	for _, char := range unfinished {
		switch {
		case inString:
			switch char {
			case '\\':
				if escaping {
					escaping = false
				} else {
					escaping = true
				}
			case '"':
				if !escaping {
					inString = false
					expectedStack = pop(expectedStack)
				}
				escaping = false
			default:
				escaping = false
			}

		case unicode.IsSpace(char):
			continue

		default:
			switch char {
			case '{':
				expectedStack = append(expectedStack, '}')
				expectingValue = false
				lastComma = false

			case '[':
				expectedStack = append(expectedStack, ']')
				expectingValue = true
				lastComma = false

			case 't':
				expectedStack = append(expectedStack, 'e', 'u', 'r')
				expectingValue = false
				lastComma = false

			case 'f':
				expectedStack = append(expectedStack, 'e', 's', 'l', 'a')
				expectingValue = false
				lastComma = false

			case 'n':
				expectedStack = append(expectedStack, 'l', 'l', 'u')
				expectingValue = false
				lastComma = false

			case ',':
				expectingValue = len(expectedStack) == 0 || expectedStack[len(expectedStack)-1] == ']'
				lastComma = true

			case '"':
				inString = true
				if !expectingValue {
					expectedStack = append(expectedStack, ':')
				}
				expectedStack = append(expectedStack, '"')
				expectingValue = false

			default:
				lastComma = false
				if len(expectedStack) > 0 && expectedStack[len(expectedStack)-1] == char {
					expectedStack = pop(expectedStack)
				}
				expectingValue = char == ':'
			}
		}
	}

	if escaping {
		sb.WriteString("\\")
	}

	if expectingValue && !lastComma {
		sb.WriteString("null")
	}

	if lastComma {
		truncated := strings.TrimSuffix(strings.TrimSpace(sb.String()), ",")
		sb.Reset()
		sb.WriteString(truncated)
	}

	if len(expectedStack) > 0 {
		sb.WriteString(strings.ReplaceAll(string(reverse(expectedStack)), ":", ":null"))
	}

	return sb.String()
}
