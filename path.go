package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	// khusus untuk file system windows, pakenya lib filepath
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))

	// filepath
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go"))   // absolute path: C:\dst
	fmt.Println(filepath.IsLocal("hello/world.go")) // realtive path
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}
