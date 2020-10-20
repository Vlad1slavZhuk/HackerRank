package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the superReducedString function below.
func superReducedString(s string) string {
	var ans string

	for i := 1; i < len(s); i++ {
		ch := string(s[i])
		ch1 := string(s[i-1])
		if ch == ch1 {
			s = s[:i-1] + s[i+1:]
			i = 0
		}
	}

	if s == "" {
		ans = "Empty String"
	} else {
		ans = s
	}

	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := superReducedString(s)

	fmt.Fprintf(writer, "%s\n", result)

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
