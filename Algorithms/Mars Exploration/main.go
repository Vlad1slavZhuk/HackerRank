package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the marsExploration function below.
func marsExploration(s string) int32 {
	var (
		ans int32
		sos string = "SOS"
	)
	count := len(s) / 3
	for i := 0; i < len(s); i++ {
		if sos[i] != s[i] {
			ans++
		}
		if count != 1 {
			sos += "SOS"
			count--
		}
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

	result := marsExploration(s)

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
