package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"uniq/uniqFunc"
)

func ReplaceFileDescriptor(fd int, file *os.File) *os.File {
	oldFile := os.NewFile(uintptr(fd), "/dev/null")
	if fd == int(os.Stdout.Fd()) {
		os.Stdout = file
	} else if fd == int(os.Stdin.Fd()) {
		os.Stdin = file
	}
	return oldFile
}

func TestOpenFile(t *testing.T) {
	_, err := OpenFile("non_existent_file.txt")
	if err == nil {
		t.Errorf("ожидается ошибка, получен nil")
	}
}

func TestMain(t *testing.T) {
	input := "a\nb\nb\na\n"
	expectedOutput := "1 a\n2 b\n1 a\n"

	reader := bytes.NewBufferString(input)
	writer := new(bytes.Buffer)

	oldStdin := os.Stdin
	oldStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdin = r
	io.WriteString(w, reader.String())
	w.Close()

	os.Stdout = ReplaceFileDescriptor(int(os.Stdout.Fd()), os.Stdout)
	options := uniqFunc.UniqOptions{Count: true}
	tmp := uniqFunc.UniqFunc(strings.Split(input, "\n"), options)
	for _, line := range tmp {
		writer.WriteString(line + "\n")
	}

	if expectedOutput != writer.String() {
		t.Errorf("ожидаемый вывод: %s, полученный: %s", expectedOutput, writer.String())
	}

	os.Stdin = oldStdin
	os.Stdout = oldStdout
}

func TestMainFileHandling(t *testing.T) {
	inputFile, err := os.CreateTemp("", "input")
	if err != nil {
		t.Fatalf("не получилось создать входной файл: %v", err)
	}
	defer os.Remove(inputFile.Name())

	content := "a\nb\nb\na\n"
	inputFile.WriteString(content)
	inputFile.Seek(0, io.SeekStart)

	outputFile, err := os.CreateTemp("", "output")
	if err != nil {
		t.Fatalf("не получилось создать выходной файл: %v", err)
	}
	defer os.Remove(outputFile.Name())

	inputFile.Close()
	outputFile.Close()

	os.Args = []string{"cmd", inputFile.Name(), outputFile.Name()}

	main()

	outContent, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("не получилось прочитать выходной файл: %v", err)
	}

	expectedOutput := "a\nb\na\n"
	if string(outContent) != expectedOutput {
		t.Errorf("ожидаемый вывод: %s, полученный: %s", expectedOutput, string(outContent))
	}
}
