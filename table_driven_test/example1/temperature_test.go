package example1

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	// Define the test cases as a table
	celsiusToFahrenheitTestCases := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{"0°C should be 32°F", 0, 32},                                                    // 0°C should be 32°F
		{"100°C should be 212°F", 100, 212},                                              // 100°C should be 212°F
		{"-40°C should be -40°F (Celsius and Fahrenheit are the same at -40)", -40, -40}, // -40°C should be -40°F (Celsius and Fahrenheit are the same at -40)
		{"37°C should be 98.6°F (normal body temperature)", 37, 98.6},                    // 37°C should be 98.6°F (normal body temperature)
		{"Absolute zero in Celsius should be -459.67°F", -273.15, -459.67},               // Absolute zero in Celsius should be -459.67°F
	}

	startTime := time.Now()
	// Iterate through the test cases and run the tests
	for _, tc := range celsiusToFahrenheitTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CelsiusToFahrenheit(tc.celsius)
			if math.Abs(result-tc.expected) > 0.01 {
				t.Errorf("CelsiusToFahrenheit(%f) = %f; want %f", tc.celsius, result, tc.expected)
			}
		})
	}

	fmt.Println("time taken", time.Now().Sub(startTime))
}

func TestFahrenheitToCelsius(t *testing.T) {
	// Define the test cases as a table
	fahrenheitToCelsiusTestCases := []struct {
		name       string
		fahrenheit float64
		expected   float64
	}{
		{"32°F should be 0°C", 32, 0},                                                    // 32°F should be 0°C
		{"212°F should be 100°C", 212, 100},                                              // 212°F should be 100°C
		{"-40°F should be -40°C (Celsius and Fahrenheit are the same at -40)", -40, -40}, // -40°F should be -40°C (Celsius and Fahrenheit are the same at -40)
		{"98.6°F should be 37°C (normal body temperature)", 98.6, 37},                    // 98.6°F should be 37°C (normal body temperature)
		{"Absolute zero in Fahrenheit should be -273.15°C", -459.67, -273.15},            // Absolute zero in Fahrenheit should be -273.15°C
	}

	// Iterate through the test cases and run the tests
	for _, tc := range fahrenheitToCelsiusTestCases {
		t.Run("", func(t *testing.T) {
			result := FahrenheitToCelsius(tc.fahrenheit)
			if math.Abs(result-tc.expected) > 0.01 {
				t.Errorf("FahrenheitToCelsius(%f) = %f; want %f", tc.fahrenheit, result, tc.expected)
			}
		})
	}
}
