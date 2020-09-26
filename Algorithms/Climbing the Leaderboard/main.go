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
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	var result = make([]int32, len(player))
	rank := make([]int32, len(ranked))
	rank[0] = 1
	for i := 1; i < len(ranked); i++ {
		if ranked[i] == ranked[i-1] {
			rank[i] = rank[i-1]
		} else {
			rank[i] = rank[i-1] + 1
		}
	}

	for i := 0; i < len(player); i++ {
		var scorePlayer = player[i]
		if scorePlayer > ranked[0] {
			result[i] = 1
		} else if scorePlayer < ranked[len(ranked)-1] {
			result[i] = rank[len(ranked)-1] + 1
		} else {
			index := binarySearch(ranked, scorePlayer)
			result[i] = rank[index]
		}
	}

	return result
}

func binarySearch(ranked []int32, key int32) int {
	var lower = 0
	var high = len(ranked) - 1

	for lower <= high {
		var mid = lower + (high-lower)/2
		if ranked[mid] == key {
			return mid
		} else if ranked[mid] < key && key < ranked[mid-1] {
			return mid
		} else if ranked[mid] > key && key >= ranked[mid+1] {
			return mid + 1
		} else if ranked[mid] < key {
			high = mid - 1
		} else if ranked[mid] > key {
			lower = mid + 1
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ranked []int32

	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		checkError(err)
		rankedItem := int32(rankedItemTemp)
		ranked = append(ranked, rankedItem)
	}

	playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		checkError(err)
		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := climbingLeaderboard(ranked, player)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

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
