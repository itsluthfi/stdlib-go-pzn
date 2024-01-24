package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "Database username")
	password := flag.String("password", "root", "Database password")
	host := flag.String("host", "localhost", "Database host")
	port := flag.Int("port", 0, "Database port")

	flag.Parse()

	// yg keprint memorynya, bukan value karena tipenya pointer
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(host)
	fmt.Println(port)

	// print valuenya
	fmt.Println(*username)
	fmt.Println(*password)
	fmt.Println(*host)
	fmt.Println(*port)

	// go run flag.go -username="luthfi izzuddin" -password=rahasia -host=421.346.153 -port=550
}
