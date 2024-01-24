package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("this is long string\nwith new line")
	reader := bufio.NewReader(input) // tujuan dari input

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line)) // dikonversi ke string karena returnnya byte
	}

	writer := bufio.NewWriter(os.Stdout) // tujuan dari output
	_, _ = writer.WriteString("Hello world\n")
	writer.WriteString("Selamat belajar\n")
	writer.Flush()
}
