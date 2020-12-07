package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the bomberMan function below.
func bomberMan(n int32, grid []string) []string {
	if n == 1 {
		return grid
	}

	if n%2 == 0 {
		for index, str := range grid {
			r := []rune(str)
			for i, ch := range r {
				if ch != 'O' {
					r[i] = 'O'
				}
			}
			grid[index] = string(r)
		}
		return grid
	}

	if n%4 == 3 {
		return Explode(grid)
	}

	return Explode(Explode(grid))
}

func Explode(grid []string) []string {
	posBomb := make([][]int, 0)

	for index, str := range grid {
		r := []rune(str)
		for i, ch := range r {
			if ch == 'O' {
				posBomb = append(posBomb, []int{index, i})
			}
		}
	}

	for index, str := range grid {
		r := []rune(str)
		for i, ch := range r {
			if ch != 'O' {
				r[i] = 'O'
			}
		}
		grid[index] = string(r)
	}

	for _, b := range posBomb {
		x, y := b[1], b[0]
		c := len(grid[0])
		r := len(grid)

		if x < c {
			grid[y] = grid[y][:x] + "." + grid[y][x+1:]
		}

		if x+1 < c {
			grid[y] = grid[y][:x+1] + "." + grid[y][x+2:]
		}
		if x-1 >= 0 {
			grid[y] = grid[y][:x-1] + "." + grid[y][x:]
		}

		if y+1 < r {
			grid[y+1] = grid[y+1][:x] + "." + grid[y+1][x+1:]
		}
		if y-1 >= 0 {
			grid[y-1] = grid[y-1][:x] + "." + grid[y-1][x+1:]
		}

	}
	return grid
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	rcn := strings.Split(readLine(reader), " ")

	rTemp, err := strconv.ParseInt(rcn[0], 10, 64)
	checkError(err)
	r := int32(rTemp)

	cTemp, err := strconv.ParseInt(rcn[1], 10, 64)
	checkError(err)
	c := int32(cTemp)
	fmt.Println(c)
	nTemp, err := strconv.ParseInt(rcn[2], 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(r); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := bomberMan(n, grid)

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
