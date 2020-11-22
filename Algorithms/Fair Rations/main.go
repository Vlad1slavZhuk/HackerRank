package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the fairRations function below.
func fairRations(B []int32) int32 {
	var (
		sum, ans int32
	)

	for _, val := range B {
		sum += val
	}

	if sum%2 != 0 {
		return -1
	}

	for i := 0; i < len(B)-1; i++ {
		if B[i]%2 != 0 {
			if i+1 < len(B) {
				B[i]++
				B[i+1]++
				ans += 2
			}
		}
	}

	if B[len(B)-1]%2 != 0 {
		ans++
	}
	return ans

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	BTemp := strings.Split(readLine(reader), " ")

	var B []int32

	for i := 0; i < int(N); i++ {
		BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
		checkError(err)
		BItem := int32(BItemTemp)
		B = append(B, BItem)
	}

	result := fairRations(B)
	if result == -1 {
		fmt.Fprintf(writer, "%s\n", "NO")
	} else {
		fmt.Fprintf(writer, "%d\n", result)
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
