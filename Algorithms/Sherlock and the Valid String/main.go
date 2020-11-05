package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

// Complete the isValid function below.
func isValid(s string) string {
	m := make(map[rune]int)

	for _, ch := range s {
		m[ch]++
	}

	mI := make([]int, len(m))
	index := 0
	for _, value := range m {
		mI[index] = value
		index++
	}
	sort.Ints(mI)

	length := len(mI)

	if length == 1 {
		return "YES"
	}
	if length == 2 && mI[0] == 1 {
		return "YES"
	}
	if mI[0] == 1 && mI[1] != mI[length-1] {
		return "NO"
	}
	if math.Abs(float64(mI[1]-mI[length-1])) > 1 {
		return "NO"
	}
	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := isValid(s)

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
