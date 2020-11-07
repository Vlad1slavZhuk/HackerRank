package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	// n - a x b поле
	// k - кол-во препятствий
	// r_q и c_q - это позиция ферзя
	// obstacles - массив препятствии
	var ans int32

	row := make([][]int32, n)
	for i := 0; i < int(n); i++ {
		row[i] = make([]int32, n)
		for j := 0; j < int(n); j++ {
			row[i][j] = 0
		}
	}
	row[r_q-1][c_q-1] = 1 // место ферзя

	// создание препятствии
	if k != 0 {
		for i := 0; i < len(obstacles); i++ {
			y := obstacles[i][0] - 1
			x := obstacles[i][1] - 1
			row[y][x] = 2
		}
	}

	ans = Left(n, r_q, c_q, row) + Right(n, r_q, c_q, row) + Up(n, r_q, c_q, row) + Down(n, r_q, c_q, row) + DiaDownLeft(n, r_q, c_q, row) + DiaUpLeft(n, r_q, c_q, row) + DiaDownRight(n, r_q, c_q, row) + DiaUpRight(n, r_q, c_q, row)
	return ans
}

func Left(n, r_q, c_q int32, arr [][]int32) int32 {
	var count int32
	if c_q == 1 {
		return 0
	}
	for x := c_q - 1; x > -1; x-- {
		if arr[r_q-1][x] == 0 {
			count++
		}
		if arr[r_q-1][x] == 2 {
			break
		}
	}
	return count
}

func Right(n, r_q, c_q int32, arr [][]int32) int32 {
	var count int32
	if c_q == n {
		return 0
	}
	for x := c_q - 1; x < n; x++ {
		if arr[r_q-1][x] == 0 {
			count++
		}
		if arr[r_q-1][x] == 2 {
			break
		}
	}
	return count
}

func Up(n, r_q, c_q int32, arr [][]int32) int32 {
	var count int32
	if r_q == 1 {
		return 0
	}
	for y := r_q - 1; y > -1; y-- {
		if arr[y][c_q-1] == 0 {
			count++
		}
		if arr[y][c_q-1] == 2 {
			break
		}
	}
	return count
}

func Down(n, r_q, c_q int32, arr [][]int32) int32 {
	var count int32
	if r_q == n {
		return 0
	}
	for y := r_q - 1; y < n; y++ {
		if arr[y][c_q-1] == 0 {
			count++
		}
		if arr[y][c_q-1] == 2 {
			break
		}
	}
	return count
}

func DiaDownLeft(n, r_q, c_q int32, arr [][]int32) int32 {
	var count int32
	if r_q == n && c_q == n {
		return 0
	}
	for y, x := r_q-1, c_q-1; x < n && y < n; x, y = x+1, y+1 {
		if arr[y][x] == 0 {
			count++
		}
		if arr[y][x] == 2 {
			break
		}
	}
	return count
}

func DiaUpLeft(n, r_q, c_q int32, arr [][]int32) int32 {
	var count int32
	if r_q == 1 && c_q == 1 {
		return 0
	}
	for y, x := r_q-1, c_q-1; x > -1 && y > -1; x, y = x-1, y-1 {
		if arr[y][x] == 0 {
			count++
		}
		if arr[y][x] == 2 {
			break
		}
	}
	return count
}

func DiaDownRight(n, r_q, c_q int32, arr [][]int32) int32 {
	var count int32
	if r_q == n && c_q == 1 {
		return 0
	}
	for y, x := r_q-1, c_q-1; x > -1 && y < n; x, y = x-1, y+1 {
		if arr[y][x] == 0 {
			count++
		}
		if arr[y][x] == 2 {
			break
		}
	}
	return count
}

func DiaUpRight(n, r_q, c_q int32, arr [][]int32) int32 {
	var count int32
	if r_q == 1 && c_q == n {
		return 0
	}
	for y, x := r_q-1, c_q-1; x < n && y > -1; x, y = x+1, y-1 {
		if arr[x][y] == 0 {
			count++
		}
		if arr[y][x] == 2 {
			break
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	r_qC_q := strings.Split(readLine(reader), " ")

	r_qTemp, err := strconv.ParseInt(r_qC_q[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(r_qC_q[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(readLine(reader), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != int(2) {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

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
