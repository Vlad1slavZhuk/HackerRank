package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the insertionSort1 function below.
func insertionSort1(n int32, arr []int32) {
	var lastV int32 = arr[n-1]

	for i := n - 1; i >= 0; i-- {
		if i-1 >= 0 && arr[i-1] > lastV {
			arr[i] = arr[i-1]
		} else {
			arr[i] = lastV
			for _, v := range arr {
				fmt.Printf("%v ", v)
			}
			fmt.Println()
			break
		}
		for _, v := range arr {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	insertionSort1(n, arr)
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
