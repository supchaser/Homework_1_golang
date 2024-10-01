package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"uniq/uniqFunc"
)

func ErrorOut(err_ error) error {
	if err_ != nil {
		fmt.Println("Ошибка", err_)
		return err_
	}
	return nil
}

func OpenFile(in string) (*os.File, error) {
	file, err := os.Open(in)
	if err != nil {
		return nil, ErrorOut(err)
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

	switch len(options.Input) {
		case 0:
			reader = os.Stdin
			output = os.Stdout
		case 1:
			var err error
			reader, err = OpenFile(options.Input[0])
			ErrorOut(err)
			output = os.Stdout
		case 2:
			var err error
			reader, err = OpenFile(options.Input[0])
			ErrorOut(err)
			resultFile, err := os.Create(options.Input[1])
			ErrorOut(err)
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
		ErrorOut(err)
	}
	writer.Flush()
}
