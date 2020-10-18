package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the camelcase function below.
func camelcase(s string) int32 {
	var count int32

	if s[0] >= 'a' && s[0] <= 'z' {
		count++
	}

	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			count++
		}
	}
	return count

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := camelcase(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
