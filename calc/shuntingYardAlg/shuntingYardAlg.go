package shuntingYardAlg

import (
	"errors"
	"fmt"
	"strconv"
)

type Token struct {
	Type  string
	Value string
}

type Stack struct {
	tokens []Token
}

func (s *Stack) IsEmpty() bool {
	return len(s.tokens) == 0
}

func (s *Stack) Push(token Token) {
	s.tokens = append(s.tokens, token)
}

func (s *Stack) Pop() (Token, bool) {
	if s.IsEmpty() {
		return Token{}, false
	}

	token := s.tokens[len(s.tokens)-1]
	s.tokens = s.tokens[:len(s.tokens)-1]

	return token, true
}

func (s *Stack) Peek() Token {
	if s.IsEmpty() {
		return Token{}
	}

	return s.tokens[len(s.tokens)-1]
}

type Queue struct {
	tokens []Token
}

func (q *Queue) IsEmpty() bool {
	return len(q.tokens) == 0
}

func (q *Queue) Enqueue(token Token) {
	q.tokens = append(q.tokens, token)
}

func (q *Queue) Dequeue() (Token, bool) {
	if q.IsEmpty() {
		return Token{}, false
	}

	token := q.tokens[0]
	q.tokens = q.tokens[1:]

	return token, true
}

func ShuntingYardAlg(tokens []Token) (float64, error) {
	var result float64
	stack := Stack{}
	queue := Queue{}

	for _, t := range tokens {
		fmt.Println(t)
	}

	for _, t := range tokens {
		switch t.Type {
		case "number":
			queue.Enqueue(t)
		case "separator":
			for !stack.IsEmpty() && stack.Peek().Type != "openingBracket" {
				token, _ := stack.Pop()
				queue.Enqueue(token)
			}
			if stack.IsEmpty() || stack.Peek().Type != "openingBracket" {
				return 0, errors.New("пропущен разделитель аргументов функции (запятая), либо пропущена открывающая скобка")
			}
		case "operator":
			for !stack.IsEmpty() && stack.Peek().Type == "operator" && getPriority(stack.Peek()) >= getPriority(t) {
				token, _ := stack.Pop()
				queue.Enqueue(token)
			}
			stack.Push(t)
		case "openingBracket":
			stack.Push(t)
		case "closingBracket":
			for !stack.IsEmpty() && stack.Peek().Type != "openingBracket" {
				token, _ := stack.Pop()
				queue.Enqueue(token)
			}
			if stack.IsEmpty() || stack.Peek().Type != "openingBracket" {
				return 0, errors.New("пропущена скобка")
			}
			stack.Pop()
		}
	}

	for !stack.IsEmpty() {
		token, _ := stack.Pop()
		if token.Type == "openingBracket" {
			return 0, errors.New("пропущена скобка")
		}
		queue.Enqueue(token)
	}

	fmt.Println("очередь")
	for _, q := range queue.tokens {
		fmt.Println(q)
	}

	for !queue.IsEmpty() {
		token, _ := queue.Dequeue()
		switch token.Type {
		case "number":
			stack.Push(token)
		case "operator":
			if token.Value == "u-" { // унарный минус
				tmp, _ := stack.Pop()
				a, _ := strconv.ParseFloat(tmp.Value, 64)
				result = -a
				stack.Push(Token{Type: "number", Value: strconv.FormatFloat(result, 'f', -1, 64)})
			} else {
				tmp1, _ := stack.Pop()
				tmp2, _ := stack.Pop()
				a, _ := strconv.ParseFloat(tmp2.Value, 64)
				b, _ := strconv.ParseFloat(tmp1.Value, 64)
				fmt.Println("A: ", a, "B: ", b)
				switch token.Value {
				case "+":
					result = a + b
				case "-":
					result = a - b
				case "*":
					result = a * b
				case "/":
					if b == 0 {
						return 0, errors.New("деление на 0")
					}
					result = a / b
				}
				stack.Push(Token{Type: "number", Value: strconv.FormatFloat(result, 'f', -1, 64)})
			}
		}
	}

	return result, nil
}
