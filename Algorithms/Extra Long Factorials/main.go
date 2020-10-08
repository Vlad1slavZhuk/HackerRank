package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// Complete the extraLongFactorials function below.
func extraLongFactorials(n int32) {
	ans := big.NewInt(1)

	for i := n; i > 0; i-- {
		nBig := big.NewInt(int64(n))
		ans = ans.Mul(ans, nBig)
	}

	fmt.Println(ans)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	extraLongFactorials(n)
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
