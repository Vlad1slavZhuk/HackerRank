package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the bigSorting function below.
func bigSorting(unsorted []string) []string {
	A, B := new(big.Int), new(big.Int)
	sort.SliceStable(unsorted, func(i, j int) bool {
		if len(unsorted[i]) < len(unsorted[j]) {
			return true
		} else if len(unsorted[i]) == len(unsorted[j]) {
			A, _ = A.SetString(unsorted[i], 10)
			B, _ = B.SetString(unsorted[j], 10)
			return A.Cmp(B) == -1
		} else {
			return false
		}
	})
	return unsorted
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

	var unsorted []string

	for i := 0; i < int(n); i++ {
		unsortedItem := readLine(reader)
		unsorted = append(unsorted, unsortedItem)
	}

	result := bigSorting(unsorted)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
