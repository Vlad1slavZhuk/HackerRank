package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the alternate function below.
func alternate(s string) int32 {
	var (
		maxL    int32
		tempSTR string
	)
	m := make(map[string]int)
	arr := make([]string, 0)
	for _, ch := range s {
		m[string(ch)]++
	}

	for v := range m {
		arr = append(arr, v)
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			tempSTR = s
			for _, ch := range s {
				if string(ch) == arr[i] || string(ch) == arr[j] {
					continue
				}
				tempSTR = strings.ReplaceAll(tempSTR, string(ch), "")
			}
			flag := true
			for k := 0; k < len(tempSTR)-1; k++ {
				if tempSTR[k] == tempSTR[k+1] {
					flag = false
					break
				}
			}
			if flag {
				if maxL < int32(len(tempSTR)) {
					maxL = int32(len(tempSTR))
				}
			}
		}
	}

	return maxL
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	lTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	l := int32(lTemp)

	s := readLine(reader)

	result := alternate(s)

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
