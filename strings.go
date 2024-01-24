package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim("   Luthfi Izzuddin   ", " "))
	fmt.Println(strings.ToLower("Luthfi Izzuddin"))
	fmt.Println(strings.ToUpper("Luthfi Izzuddin"))
	fmt.Println(strings.Split("Luthfi Izzuddin", " "))
	fmt.Println(strings.Contains("Luthfi Izzuddin", "uddin"))
	fmt.Println(strings.ReplaceAll("Luthfi Izzuddin Luthfi Hanif", "Luthfi", "Lupi"))
	fmt.Println(strings.Replace("Luthfi Izzuddin Luthfi Hanif", "Luthfi", "Lupi", 1))
}
