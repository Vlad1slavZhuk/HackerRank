package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"math"
)

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int32) int32 {
	a := [][][]int32{
		{
			{4, 9, 2},
			{3, 5, 7},
			{8, 1, 6},
		},
		{
			{8, 1, 6},
			{3, 5, 7},
			{4, 9, 2},
		},
		{
			{2, 9, 4},
			{7, 5, 3},
			{6, 1, 8},
		},
		{
			{6, 1, 8},
			{7, 5, 3},
			{2, 9, 4},
		},
		{
			{8, 3, 4},
			{1, 5, 9},
			{6, 7, 2},
		},
		{
			{6, 7, 2},
			{1, 5, 9},
			{8, 3, 4},
		},
		{
			{4, 3, 8},
			{9, 5, 1},
			{2, 7, 6},
		},
		{
			{2, 7, 6},
			{9, 5, 1},
			{4, 3, 8},
		},
	}

	min := 100000
	var sum float64
	for i := 0; i < 8; i++ {
		sum = 0
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				sum += math.Abs(float64(a[i][j][k] - s[j][k]))
			}
		}
		if min > int(sum) {
			min = int(sum)
		}
	}

	return int32(min)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(readLine(reader), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

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
