package main

import (
	"fmt"
)

func GalToPound(gal float32) string {
	var msg string
	var refPound float32 = 6.8

	if gal > 0 {
		var valTot = refPound * gal
		msg = "Total of Pounds is: " + fmt.Sprint(valTot) + " lb"
	} else {
		msg = "Gallons must be value."
	}

	return msg
}

func main() {

	var gal float32

	fmt.Print("Enter the number of gallons (Ex.: 4800.50 or 7000): ")
	fmt.Scanln(&gal)

	fmt.Println(GalToPound(gal))
}
