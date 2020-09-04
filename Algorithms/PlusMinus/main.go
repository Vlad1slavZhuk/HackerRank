package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	var pos, neg, zero int
	var posF, negF, zeroF float64

	for i := 0; i < len(arr); i++ {
		if arr[i] > 100 || arr[i] < -100 {
			break
		} else {
			if arr[i] > 0 {
				pos++
			} else if arr[i] < 0 {
				neg++
			} else {
				zero++
			}
		}
	}

	posF = float64(pos) / float64(len(arr))
	negF = float64(neg) / float64(len(arr))
	zeroF = float64(zero) / float64(len(arr))

	fmt.Printf("%f\n%f\n%f\n", posF, negF, zeroF)
}

//----------------------------------------------------------
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	//----------------------------------------------------------
	if n < 0 || n > 100 {
		return
	}
	//----------------------------------------------------------
	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
