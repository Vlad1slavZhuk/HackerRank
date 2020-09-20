package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the pageCount function below.
 */
func pageCount(n int32, p int32) int32 {
	/*
	 * Write your code here.
	 */
	var count1, count2, result int32

	if p == 1 || p == n {
		count1 = 0
	}

	for i := 1; i < int(n); i += 2 {
		if i == int(p) || i-1 == int(p) {
			break
		}
		count1++
	}

	if n%2 != 0 {
		for i := n; i > 0; i -= 2 {
			if i == p || i-1 == p {
				break
			}
			count2++
		}
	} else {
		for i := n + 1; i > 0; i -= 2 {
			if i == p || i-1 == p {
				break
			}
			count2++
		}
	}

	if count1 < count2 {
		result = count1
	} else {
		result = count2
	}

	return result
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

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int32(pTemp)

	result := pageCount(n, p)

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
