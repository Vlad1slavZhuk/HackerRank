package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the pangrams function below.
func pangrams(s string) string {
	var ans string = "not pangram"
	m := make(map[string]int)

	for _, ch := range s {
		if ch == ' ' {
			continue
		}
		c := strings.ToLower(string(ch))
		m[c]++
	}

	if len(m) >= 26 {
		ans = "pangram"
	}

	return ans

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := pangrams(s)

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
