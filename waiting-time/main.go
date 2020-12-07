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
 * Complete the 'waitingTime' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY tickets
 *  2. INTEGER p
 */

func waitingTime(tickets []int32, p int32) int64 {
	// Write your code here
	var i int32 = 0
	n := int32(len(tickets)) - 1
	var t int64
	for true {
		if i > n {
			i = 0
		}
		if i != p && tickets[i] == 0 {
			i++
			continue
		}
		tickets[i]--
		t++
		if i == p && tickets[i] == 0 {
			break
		}
		i++
	}
	return t
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	ticketsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var tickets []int32

	for i := 0; i < int(ticketsCount); i++ {
		ticketsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		ticketsItem := int32(ticketsItemTemp)
		tickets = append(tickets, ticketsItem)
	}

	pTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	p := int32(pTemp)

	result := waitingTime(tickets, p)

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

