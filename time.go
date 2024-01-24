package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now) // waktunya sama kayak local

	var utc time.Time = time.Date(2023, time.September, 25, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local()) // konversi dari utc ke localtime

	formatter := "2006-Jan-02 15:04:05" // formatnya ngikutin dokumentasi, dibaca dulu dokumentasinya
	valueToFormat := "2023-Sep-25 23:15:00"

	valueTime, err := time.Parse(formatter, valueToFormat)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(valueTime)

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day()) // day of the month, gabisa kalo nama hari, returnnya int
	fmt.Println(valueTime.Hour())

	// duration
	var duration1 time.Duration = 100 * time.Minute
	duration2 := 10 * time.Millisecond
	duration3 := duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("Duration %d\n", duration3) // outputnya dalam format nanosecond
}
