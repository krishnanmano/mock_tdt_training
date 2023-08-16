package example1

// CelsiusToFahrenheit converts Celsius to Fahrenheit.
func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

// FahrenheitToCelsius converts Fahrenheit to Celsius.
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
