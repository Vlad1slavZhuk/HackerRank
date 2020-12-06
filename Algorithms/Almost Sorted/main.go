package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the almostSorted function below.
func almostSorted(arr []int32) {
	var right, left int

	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			right = i
			break
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			left = i
			break
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	x := true

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			x = false
		}
	}

	if x {
		fmt.Println("yes")
		fmt.Println("swap", left+1, right+1)
	} else {
		for i, j := left+1, right-1; i <= j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}

		x = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				x = false
			}
		}
		if x {
			fmt.Println("yes")
			fmt.Println("reverse", left+1, right+1)
		} else {
			fmt.Println("no")
		}
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

	almostSorted(arr)
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
