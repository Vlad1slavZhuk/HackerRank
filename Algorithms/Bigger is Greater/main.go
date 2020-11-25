package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the biggerIsGreater function below.
func biggerIsGreater(w string) string {
	var (
		ans string = "no answer"
	)
	a := make([]int32, len(w))
	w = strings.TrimSpace(w)

	for pos, ch := range w {
		if ch >= 'a' && ch <= 'z' {
			a[pos] = ch
		}
	}

	i := len(a) - 1
	for {
		if i <= 0 || a[i-1] < a[i] {
			break
		}
		i--
	}
	if i <= 0 {
		return ans
	}

	j := len(a) - 1
	for {
		if i == j {
			break
		}
		if a[j] > a[i-1] {
			break
		}
		j--
	}
	a[i-1], a[j] = a[j], a[i-1]

	j = len(a) - 1
	for {
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
		j--
		i++
	}
	ans = ""
	for i := 0; i < len(a); i++ {
		ans += string(a[i])
	}
	return ans

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	TTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		w := readLine(reader)

		result := biggerIsGreater(w)

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
