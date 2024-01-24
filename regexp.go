package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`u([a-z])i`) // baca lengkapnya di dokumentasi regexp & re2-google

	fmt.Println(regex.MatchString("upi"))
	fmt.Println(regex.MatchString("uGi"))
	fmt.Println(regex.MatchString("uji"))
	fmt.Println(regex.MatchString("u3i"))

	fmt.Println(regex.FindAllString("upi uGi uji u3i ubi", 10))
}
