package main

import (
	"flag"
	"fmt"
	"funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahrenheit float64
var celcius float64
var kelvin float64
var out string
var funfacts string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahrenheit, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celcius, "C", 0.0, "temperatur i grader celcius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")

}
func main() {

	flag.Parse()

	/**
	  fmt.Println("len(flag.Args())", len(flag.Args()))
	  fmt.Println("flag.NFlag()", flag.NFlag())

	  fmt.Println(isFlagPassed("out"))
	*/

	//Fahrenheit til Celcius
	if out == "C" && isFlagPassed("F") {
		fmt.Printf("%v°F er = %.2f°C", fahrenheit, conv.FarhenheitToCelsius(fahrenheit))
	}

	//Fahrenheit til Kelvin
	if out == "K" && isFlagPassed("F") {
		fmt.Printf("%v°F er = %.2f°K", fahrenheit, conv.FarhenheitToKelvin(fahrenheit))
	}

	//Celcius til Kelvin
	if out == "K" && isFlagPassed("C") {
		fmt.Printf("%v°F er = %.2f°C", celcius, conv.CelsiusToKelvin(celcius))
	}

	//Celcius til Fahrenheit
	if out == "F" && isFlagPassed("C") {
		fmt.Printf("%v°F er = %.2f°C", celcius, conv.CelsiusToFahrenheit(celcius))
	}

	//Kelvin til Celcius
	if out == "C" && isFlagPassed("K") {
		fmt.Printf("%v°K er = %.2f°C", kelvin, conv.KelvintoCelcius(kelvin))
	}

	//Kelvin til Fahrenheit
	if out == "F" && isFlagPassed("K") {
		fmt.Printf("%v°K er = %.2f°F", kelvin, conv.KelvinToFarhenheit(kelvin))
	}
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
