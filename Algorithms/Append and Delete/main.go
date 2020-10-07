package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) string {
	var result, minStr string

	if len(s) <= len(t) {
		minStr = s
	} else {
		minStr = t
	}

	count := 0

	for i := 0; i < len(minStr); i++ {
		if s[i] != t[i] {
			break
		}
		count++
	}

	diff := int32((len(s) - count) + (len(t) - count))
	totalLen := int32(len(s) + len(t))

	switch {
	case diff > k:
		result = "No"
	case diff == k:
		result = "Yes"
	case totalLen%2 == k%2:
		result = "Yes"
	case totalLen < k:
		result = "Yes"
	default:
		result = "No"
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	t := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := appendAndDelete(s, t, k)

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
