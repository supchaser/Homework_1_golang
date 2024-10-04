package uniqFunc

import (
	"bufio"
	"strconv"
	"strings"
)

// Структура для хранения часто используемых данных
type MyStruct struct {
	prev string
	result []string
	count int
	flag bool
}

// Функция для подсчета повторяющихся строк
func CountFunc (input *bufio.Scanner) []string {
	values := MyStruct{count: 1}
	for input.Scan() {
		txt := input.Text()
		if txt == values.prev {
			values.count++
		} else {
			if values.flag {
				values.result = append(values.result, strconv.Itoa(values.count) + " " + values.prev)
			}
			values.prev = txt
			values.count = 1
		}

		values.flag = true
	}

	if values.prev != "" {
		values.result = append(values.result, strconv.Itoa(values.count) + " " + values.prev)
	}

	return values.result
}

// Функция для вывода только повторяющихся строк
func RepeatedFunc (input *bufio.Scanner) []string {
	values := MyStruct{count: 1}
	for input.Scan() {
		txt := input.Text()
		if txt == values.prev {
			values.count++
		} else {
			if values.flag  && values.count > 1{
				values.result = append(values.result, values.prev)
			}
			values.prev = txt
			values.count = 1
		}
		
		values.flag = true
	}

	if values.flag && values.count > 1 {
		values.result = append(values.result, values.prev)
	}

	return values.result
}

// Функция для вывода только уникальных строк
func UniqueFunc (input *bufio.Scanner) []string {
	values := MyStruct{count: 1}
	for input.Scan() {
		txt := input.Text()
		if txt == values.prev {
			values.count++
		} else {
			if values.flag  && values.count == 1{
				values.result = append(values.result, values.prev)
			}
			values.prev = txt
			values.count = 1
		}

		values.flag = true
	}

	if values.flag && values.count == 1 {
		values.result = append(values.result, values.prev)
	}

	return values.result
}

// Функция для игнорирования регистра
func IgnoreCaseFunc (input *bufio.Scanner) []string {
	values := MyStruct{}
	for input.Scan() {
		txt := input.Text()
		toLower := strings.ToLower(txt)

		if toLower == values.prev {
			continue
		}

		values.prev = toLower
		values.result = append(values.result, txt)
	}

	return values.result
}

// Функция для пропуска определенного количества полей
func SkipFieldsFunc (input *bufio.Scanner, valueSkipFields int) []string {
	values := MyStruct{}
	for input.Scan() {
		txt := input.Text()
		fields := strings.Fields(txt)
		if valueSkipFields > len(fields) {
			values.result = append(values.result, txt)
			continue
		}

		cutString := strings.Join(fields[valueSkipFields:], " ")
		if cutString == values.prev {
			continue
		}

		values.prev = cutString
		values.result = append(values.result, txt)
	}

	return values.result
}

// Функция для пропуска определенного количества символов
func SkipCharsFunc (input *bufio.Scanner, valueSkipChars int) []string {
	values := MyStruct{}
	for input.Scan() {
		txt := input.Text()
		if valueSkipChars > len(txt) {
			values.result = append(values.result, txt)
			continue
		}

		cutString := txt[valueSkipChars:]
		if cutString == values.prev {
			continue
		}

		values.prev = cutString
		values.result = append(values.result, txt)
	}

	return values.result
}

// Функция по умолчанию
func DefaultFunc (input *bufio.Scanner) []string {
	values := MyStruct{}
	for input.Scan() {
		txt := input.Text()
		if txt == values.prev {
			continue
		}

		values.prev = txt
		values.result = append(values.result, txt)
	}

	return values.result
}

func UniqFunc(input []string, options UniqOptions) []string {
	reader := strings.NewReader(strings.Join(input, "\n"))
	in := bufio.NewScanner(reader)
	var result []string
	
	switch {
	case options.Count:
		result = CountFunc(in)
	case options.Repeated:
		result = RepeatedFunc(in)
	case options.Unique:
		result = UniqueFunc(in)
	case options.IgnoreCase:
		result = IgnoreCaseFunc(in)
	case options.SkipFields != 0:
		result = SkipFieldsFunc(in, options.SkipFields)
		if options.SkipChars != 0 {
			reader_skip := strings.NewReader(strings.Join(result, "\n"))
			in_skip := bufio.NewScanner(reader_skip)
			result = SkipCharsFunc(in_skip, options.SkipChars)
		}
	case options.SkipChars != 0:
		result = SkipCharsFunc(in, options.SkipChars)
	default:
		result = DefaultFunc(in)
	}

	return result
}
