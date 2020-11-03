package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the funnyString function below.
func funnyString(s string) string {
	var (
		length int   = len(s)
		a1     []int = make([]int, length)
		a2     []int = make([]int, length)
	)

	for i := 0; i < length; i++ {
		a1[i] = int(s[i])
	}

	index := 0
	for i := length - 1; i >= 0; i-- {
		a2[index] = int(s[i])
		index++
	}

	for i := 0; i < length-1; i++ {
		l1 := math.Abs(float64(a1[i] - a1[i+1]))
		l2 := math.Abs(float64(a2[i] - a2[i+1]))

		if l1 != l2 {
			return "Not Funny"
		} else {
			continue
		}
	}

	return "Funny"

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := funnyString(s)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
