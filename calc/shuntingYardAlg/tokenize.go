package shuntingYardAlg

import (
	"unicode"
)

// Разбиение на токены
func Tokenize(input string) []Token {
	tokens := []Token{}
	var current string
	var havePoint bool

	for _, r := range input {
		switch {
		case r == ' ':
			if current != "" {
				tokens = append(tokens, Token{Type: "number", Value: current})
				current = ""
				havePoint = false
			}
		case r == '+' || r == '-' || r == '*' || r == '/':
			if current != "" {
				tokens = append(tokens, Token{Type: "number", Value: current})
				current = ""
				havePoint = false
			}

			if r == '-' && (len(tokens) == 0 || tokens[len(tokens)-1].Type == "operator" || tokens[len(tokens)-1].Type == "openingBracket") {
				tokens = append(tokens, Token{Type: "operator", Value: "u-"})
			} else {
				tokens = append(tokens, Token{Type: "operator", Value: string(r)})
			}
		case r == '(':
			if current != "" {
				tokens = append(tokens, Token{Type: "number", Value: current})
				current = ""
				havePoint = false
			}

			if len(tokens) > 0 && (tokens[len(tokens)-1].Type == "number" || tokens[len(tokens)-1].Type == "closingBracket") {
				tokens = append(tokens, Token{Type: "operator", Value: "*"})
			}

			tokens = append(tokens, Token{Type: "openingBracket", Value: string(r)})
		case r == ')':
			if current != "" {
				tokens = append(tokens, Token{Type: "number", Value: current})
				current = ""
				havePoint = false
			}

			tokens = append(tokens, Token{Type: "closingBracket", Value: string(r)})
		case r == ',':
			if current != "" {
				tokens = append(tokens, Token{Type: "number", Value: current})
				current = ""
				havePoint = false
			}

			tokens = append(tokens, Token{Type: "separator", Value: string(r)})
		case r == '.':
			if havePoint {
				continue
			}
			
			if current == "" {
				current = "0"
			}

			current += string(r)
			havePoint = true
		case unicode.IsDigit(r):
			current += string(r)
		default:
			continue
		}
	}

	if current != "" {
		tokens = append(tokens, Token{Type: "number", Value: current})
	}

	return tokens
}

func getPriority(token Token) int {
	switch token.Value {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "u-":
		return 3
	default:
		return 0
	}
}
