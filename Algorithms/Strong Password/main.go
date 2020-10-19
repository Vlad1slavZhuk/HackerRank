package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Complete the minimumNumber function below.
func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	var count int32
	Digits, _ := regexp.Compile(`\d`)
	UpperAlp, _ := regexp.Compile(`[A-Z]`)
	LowerAlp, _ := regexp.Compile(`[a-z]`)
	Special, _ := regexp.Compile(`\W`)

	if !Digits.Match([]byte(password)) {
		count++
	}
	if !UpperAlp.Match([]byte(password)) {
		count++
	}
	if !LowerAlp.Match([]byte(password)) {
		count++
	}
	if !Special.Match([]byte(password)) {
		count++
	}

	if count+int32(len(password)) < 6 {
		count += 6 - (count + int32(len(password)))
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	password := readLine(reader)

	answer := minimumNumber(n, password)

	fmt.Fprintf(writer, "%d\n", answer)

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
