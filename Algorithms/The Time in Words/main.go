package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the timeInWords function below.
func timeInWords(h int32, m int32) string {
	var ans string
	hoursArr := map[int32]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
	}
	minArr := map[int32]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		21: "twenty one",
		22: "twenty two",
		23: "twenty three",
		24: "twenty four",
		25: "twenty five",
		26: "twenty six",
		27: "twenty seven",
		28: "twenty eight",
		29: "twenty nine",
	}
	if m == 0 {
		ans = fmt.Sprintf("%v o' clock", hoursArr[h])
	} else {
		if m == 1 {
			ans = fmt.Sprintf("%v minute past %v", minArr[m], hoursArr[h])
			return ans
		} else if m > 1 && m < 31 {
			if m == 15 {
				ans = fmt.Sprintf("quarter past %v", hoursArr[h])
				return ans
			} else if m == 30 {
				ans = fmt.Sprintf("half past %v", hoursArr[h])
				return ans
			}
			ans = fmt.Sprintf("%v minutes past %v", minArr[m], hoursArr[h])
			return ans
		} else if m > 30 && m < 60 {
			hours := h + 1
			if m == 45 {
				ans = fmt.Sprintf("quarter to %v", hoursArr[hours])
				return ans
			}
			minutes := 60 - m
			ans = fmt.Sprintf("%v minutes to %v", minArr[minutes], hoursArr[hours])
			return ans
		}
	}
	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	hTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	h := int32(hTemp)

	mTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := timeInWords(h, m)

	fmt.Fprintf(writer, "%s\n", result)

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
