// Package weather provides tools to forecast weather.
package weather

var (
	// CurrentCondition represents a current weather condition.
	CurrentCondition string
	// CurrentLocation represents a current location.
	CurrentLocation string
)

// Forecast returns a string value equal to current weather condition for specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
