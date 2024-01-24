package main

import (
	"errors"
	"fmt"
)

//* best practicenya kalo bikin error didefinisiin pake multivars
// bisa pake const, tapi kalo pake const gabisa panggil fungsi, ex: errors.New()

var (
	ValidationError = errors.New("Data tidak tervalidasi")
	NotFoundError   = errors.New("Data tidak ditemukan")
)

func getByID(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "Luthfi" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := getByID("Luthf")

	if err != nil {
		if errors.Is(err, ValidationError) { // mengecek jenis error
			fmt.Println(ValidationError.Error())
		} else if errors.Is(err, NotFoundError) {
			fmt.Println(NotFoundError.Error())
		} else {
			fmt.Println("Unknown error", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
