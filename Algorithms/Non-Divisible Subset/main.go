package main

import (
	"bufio"
	"fmt"
	"io"

	//"os"
	//"strconv"
	"strings"
)

/*
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */

func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	var (
		ans     int32
		remains = make([]int32, k)
	)

	for i := 0; i < len(s); i++ {
		remains[s[i]%k]++
	}

	if remains[0] > 0 {
		ans++
	}

	for i := 1; i < len(remains); i++ {
		if int32(i) == k-int32(i) {
			ans++
		} else {
			if remains[i] >= remains[int(k)-i] {
				ans += remains[i]
			} else {
				ans += remains[int(k)-i]
			}
			remains[i] = 0
			remains[int(k)-i] = 0
		}
	}

	return ans
}

func main() {
	/*reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}*/

	k, s := int32(3), []int32{1, 7, 2, 4}
	result := nonDivisibleSubset(k, s)
	fmt.Println(result)

	/*fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()*/
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
