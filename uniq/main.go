package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"uniq/uniqFunc"
)

func OpenFile(in string) (*os.File, error) {
	file, err := os.Open(in)
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return file, err
	}
	return file, nil
}

func main() {
	options := uniqFunc.GetFlags()
	err := uniqFunc.Validation(options)
	if err != nil {
		fmt.Println("Ошибка валидации: ", err)
		return
	}

	var tmp []string
	var reader io.Reader
	var output io.Writer

	// определяем входной и выходной потоки в зависимости от длины input'а
	switch len(options.Input) {
		case 0:
			reader = os.Stdin
			output = os.Stdout
		case 1:
			var err error
			reader, err = OpenFile(options.Input[0])
			if err != nil {
				fmt.Println("ошибка открытия файла", err)
				return
			}
			output = os.Stdout
		case 2:
			var err error
			reader, err = OpenFile(options.Input[0])
			if err != nil {
				fmt.Println("ошибка открытия файла", err)
				return
			}
			resultFile, err := os.Create(options.Input[1])
			if err != nil {
				fmt.Println("ошибка создания файла", err)
				return
			}
			defer resultFile.Close()
			output = resultFile
		}

	in := bufio.NewScanner(reader)
	for in.Scan() {
		txt := in.Text()
		tmp = append(tmp, txt)
	}

	writer := bufio.NewWriter(output)
	result := uniqFunc.UniqFunc(tmp, options)

	for i := 0; i < len(result); i++ {
		line := result[i]
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("ошибка записи строки", err)
			return
		}
	}
	writer.Flush()
}
