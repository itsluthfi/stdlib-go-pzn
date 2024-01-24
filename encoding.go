package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("Luthfi Izzuddin Hanif")) // dikonversi ke byte dulu
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(decoded)         // ini kalo belum dikonversi ke string
	fmt.Println(string(decoded)) // harus dikonversi ke string dulu

	// baca csv
	csvString := "luthfi,izzuddin,hanif\n" + "ima,zumi,ito\n" + "za,na,van\n"

	reader := csv.NewReader(strings.NewReader(csvString)) // pake strings karena csvnya hard-coded
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	// tulis csv
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"luthfi", "izzuddin", "hanif"})
	_ = writer.Write([]string{"ima", "zumi", "ito"})
	_ = writer.Write([]string{"za", "na", "van"})

	writer.Flush()
}
