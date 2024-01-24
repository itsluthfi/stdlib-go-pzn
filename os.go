package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args // buat ngeprint argument yg ada di perintah console

	for _, arg := range args {
		fmt.Println(arg)
	}

	// mendapatkan hostname
	host, err := os.Hostname()
	if err == nil {
		fmt.Println(host)
	} else {
		fmt.Println(err.Error())
	}
}
