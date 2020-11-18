package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the happyLadybugs function below.
func happyLadybugs(b string) string {
	chars := make(map[rune]int)
	charsR := make([]byte, 0)

	for _, ch := range b {
		chars[ch]++
	}

	if _, ok := chars['_']; !ok {
		lm := len(chars)
		for _, val := range chars {
			if val == 1 {
				return "NO"
			}
		}
		if lm == 1 {
			return "YES"
		}
		i := -1
		for k := 0; k < len(b); k++ {
			if !Contains(charsR, b[k]) {
				charsR = append(charsR, b[k])
				i++
			} else {
				if b[k] != charsR[i] {
					return "NO"
				}
			}
		}
		return "YES"
	}

	for key, val := range chars {
		if val == 1 && key != '_' {
			return "NO"
		}
	}

	return "YES"
}

func Contains(a []byte, x byte) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	gTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	g := int32(gTemp)

	for gItr := 0; gItr < int(g); gItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)
		fmt.Println(n)

		b := readLine(reader)

		result := happyLadybugs(b)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
