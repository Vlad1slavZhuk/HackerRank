package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	var min, max, sum int64

	for i := 0; i < len(arr); i++ {
		sum = 0
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			sum += int64(arr[j])
		}
		if min == 0 && max == 0 {
			min, max = sum, sum
		}
		if sum > max {
			max = sum
		}
		if sum < min {
			min = sum
		}
	}
	if min > max {
		min, max = max, min
	}
	fmt.Println(min, max)
}

//-------------------------------------------------------------
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
