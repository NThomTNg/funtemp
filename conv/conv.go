package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

//func FarhenheitToCelsius(value float64) float64 {
// Her skal du implementere funksjonen
// Du skal ikke formattere float64 i denne funksjonen
// Gjør formattering i main.go med fmt.Printf eller
// lignende
//return 56.67

// De andre konverteringsfunksjonene implementere her

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(farenheit float64) float64 {
	return (farenheit - 32) * 5 / 9
}

// Konverterere Celsius til Farhenheit
func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

//Konverterer Kelvin til Celsius
func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
