package main

import (
	"fmt"
)

var refOneGalInPound float32 = 6.8 // On Gal is equals 6.8 pounds (In simulator with the Zibo)
var refPoundToKg float32 = 2.205   // Tax to calculate between Kile and pounds

func GalToPound(gal float32) string {
	var msg string

	if gal > 0 {
		var valTot = refOneGalInPound * gal
		msg = "Total of Pounds is: " + fmt.Sprint(valTot) + " lb"
	} else {
		msg = "Gallons must be value."
	}

	return msg
}

func GalToKilo(gal float32) string {
	var msg string

	if gal > 0 {

		var valTot = refOneGalInPound * gal
		valTot = valTot / refPoundToKg
		msg = "Total of Kile is: " + fmt.Sprint(valTot) + " kg"
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
	fmt.Println(GalToKilo(gal))
}
