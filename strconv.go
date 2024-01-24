package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	result1, err1 := strconv.ParseFloat("315.134", 0)
	if err1 != nil {
		fmt.Println(err1.Error())
	} else {
		fmt.Println(result1)
	}

	result2, err2 := strconv.ParseInt("531", 10, 0)
	if err2 != nil {
		fmt.Println(err2.Error())
	} else {
		fmt.Println(result2)
	}

	result3, err3 := strconv.Atoi("124")
	if err3 != nil {
		fmt.Println(err3.Error())
	} else {
		fmt.Println(result3)
	}

	result4 := strconv.Itoa(125)
	fmt.Println(result4)

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)
}
