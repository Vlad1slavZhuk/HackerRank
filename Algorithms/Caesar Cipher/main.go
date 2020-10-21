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

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	var (
		str string
	)

	for _, v := range s {
		ch := string(v)

		if ok, _ := regexp.MatchString("[a-zA-Z]", ch); ok {
			if ok, _ := regexp.MatchString("[A-Z]", ch); ok {
				caesar := (((v - 'A') + k) % 26) + 'A'
				str += string(caesar)
			}

			if ok, _ := regexp.MatchString("[a-z]", ch); ok {
				caesar := (((v - 'a') + k) % 26) + 'a'
				str += string(caesar)
			}
		} else {
			str += ch
		}

	}
	return str

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

	s := readLine(reader)

	if n != int32(len(s)) {
		return
	}

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

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
