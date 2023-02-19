package main

import (
	"flag"

	"github.com/SSneakySnek/funtemp/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var out string
var funfacts string
var C float64
var F float64
var K float64

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	//Definer flagg-variablene
	flag.Float64Var(&C, "C", 0.0, "temperatur i grader i celsius")
	flag.Float64Var(&F, "F", 0.0, "temperatur i grader i farenheit")
	flag.Float64Var(&K, "K", 0.0, "temperatur i grader i kelvin")

	flag.StringVar(&out, "out", "", "temperaturskala for resultat")

}

func main() {
	flag.Parse()
	var res float64
	//standardverdi 0
	if C != 0 {
		if out == "F" {
			res = conv.CelsiusToFarenheit(C)
		} else if out == "K" {
			res = C + 273.15
		}
	}

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.
	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.
	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())
	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	//fmt.Println(fahr, out, funfacts)
	//fmt.Println("len(flag.Args())", len(flag.Args()))
	//fmt.Println("flag.NFlag()", flag.NFlag())
	//fmt.Println(isFlagPassed("out"))
	// Eksempel på enkel logikk
	//if out == "C" && isFlagPassed("F") {
	// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
	// skal returnere °C
	//fmt.Println("0°F er -17.78°C")
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
