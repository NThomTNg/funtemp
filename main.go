package main

import (
	"flag"
	"fmt"
)

var C float64
var F float64
var K float64
var toScale string

func main() {
	var (
		celsius    float64
		fahrenheit float64
		kelvin     float64
	)

	flag.Float64Var(&celsius, "c", 0.0, "Temperatur i Celsius")
	flag.Float64Var(&fahrenheit, "f", 0.0, "Temperatur i Fahrenheit")
	flag.Float64Var(&kelvin, "k", 0.0, "Temperatur i Kelvin")
	flag.StringVar(&toScale, "to", "", "Temperaturskala for reslutatet (c/f/k)")

	flag.Parse()

	if toScale == "" {
		fmt.Println("Spesifiser en temperatur skala for å konverter til å bruke -to flag (c/f/k)")
		return
	}

	if (celsius != 0 && (fahrenheit != 0 || kelvin != 0)) ||
		(fahrenheit != 0 && (celsius != 0 || kelvin != 0)) ||
		(kelvin != 0 && (celsius != 0 || fahrenheit != 0)) {
		fmt.Println("En temperatur verdi for å konverter")
		return
	}

	switch toScale {
	case "c":
		if celsius != 0 {
			fmt.Printf("%.2f°C\n", celsius)
		} else if fahrenheit != 0 {
			celsius = (fahrenheit - 32) * 5 / 9
			fmt.Printf("%.2f°F is %.2f°C\n", fahrenheit, celsius)
		} else {
			celsius = kelvin - 273.15
			fmt.Printf("%.2fK is %.2f°C\n", kelvin, celsius)
		}
	case "f":
		if fahrenheit != 0 {
			fmt.Printf("%.2f°F\n", fahrenheit)
		} else if celsius != 0 {
			fahrenheit = celsius*9/5 + 32
			fmt.Printf("%.2f°C is %.2f°F\n", celsius, fahrenheit)
		} else {
			fahrenheit = kelvin*9/5 - 459.67
			fmt.Printf("%.2fK is %.2f°F\n", kelvin, fahrenheit)
		}
	case "k":
		if kelvin != 0 {
			fmt.Printf("%.2fK\n", kelvin)
		} else if celsius != 0 {
			kelvin = celsius + 273.15
			fmt.Printf("%.2f°C is %.2fK\n", celsius, kelvin)
		} else {
			kelvin = (fahrenheit + 459.67) * 5 / 9
			fmt.Printf("%.2f°F is %.2fK\n", fahrenheit, kelvin)
		}
	default:
		fmt.Printf("Ukjent temperatur skala %s\n", toScale)
	}
}
