package main

import (
	"fmt"
)

var refOneGalInPound float32 = 6.8 // On Gal is equals 6.8 pounds (In simulator with the Zibo)
var refPoundToKg float32 = 2.205   // Tax to calculate between Kile and pounds

func GalToPound(gal float32) string {
	var msg string

	var valTot = refOneGalInPound * gal
	msg = "Total of Gallons to Pounds is: " + fmt.Sprint(valTot) + " lb"

	return msg
}

func GalToKilo(gal float32) string {
	var msg string

	var valTot = refOneGalInPound * gal
	valTot = valTot / refPoundToKg
	msg = "Total of Gallons to Kile is: " + fmt.Sprint(valTot) + " kg"

	return msg
}

func KileToGal(kg float32) string {
	var msg string

	var valTot = kg / refOneGalInPound
	valTot = valTot * refPoundToKg
	msg = "Total of Kile to Gallons is: " + fmt.Sprint(valTot) + " gal"

	return msg
}

func PoundToGal(pound float32) string {
	var msg string

	var valTot = pound / refOneGalInPound
	msg = "Total of Pounds to Gallons is: " + fmt.Sprint(valTot) + " gal"

	return msg
}

func main() {

	var gal, pounds, kg float32

	fmt.Print("Enter the number of gallons (Ex.: 4800.50 or 7000): ")
	fmt.Scanln(&gal)

	fmt.Print("Enter the number of pounds (Ex.: 4800.50 or 7000): ")
	fmt.Scanln(&pounds)

	fmt.Print("Enter the number of kile (Ex.: 4800.50 or 7000): ")
	fmt.Scanln(&kg)

	fmt.Println("")

	if gal > 0 {
		fmt.Println(GalToPound(gal))
		fmt.Println(GalToKilo(gal))
	} else if gal == 0 {
		fmt.Println("0 gal")
	} else {
		fmt.Println("Gallons must be value.")
	}

	if pounds > 0 {
		fmt.Println(PoundToGal(pounds))
	} else if pounds == 0 {
		fmt.Println("0 lb")
	} else {
		fmt.Println("Pounds must be value.")
	}

	if kg > 0 {
		fmt.Println(KileToGal(kg))
	} else if kg == 0 {
		fmt.Println("0 kg")
	} else {
		fmt.Println("Kile must be value.")
	}
}
