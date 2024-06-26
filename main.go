package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

/*
 * Complete the 'camelcase' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func camelcase(s string) int32 {
	var cw int32
	if len(s) == 0 {
		return cw
	}
	for i, c := range s {
		if i == 0 {
			cw++
			continue
		}
		if unicode.IsUpper(c) {
			cw++
		}
	}
	return cw
}

func caesarCipher(s string, k int32) string {
	if k < 0 || k > 100 {
		return 0
	}
	var encrypted string
	var shifted int
	asciiZ := int('z')
	asciiA := int('a')
	for _, c := range s {
		if !unicode.IsLetter(c) {
			encrypted += string(c)
			continue
		}
		ascii := int(c)
		if ascii+int(k) > asciiZ {
			shifted = asciiA + (ascii + int(k) - asciiZ) - 1
		} else {
			shifted = ascii + int(k)
		}
		encrypted += string(shifted)
	}
	return encrypted
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer stdout.Close()
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	s := readLine(reader)
	result := camelcase(s)
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
