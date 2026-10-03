/* Package weather includes data and methods
related to weather forecasting. */
package weather

var (
    // CurrentCondition stores the current condition.
	CurrentCondition string
    // CurrentLocation stores the current location.
	CurrentLocation  string
)

// Forecast gets the forecast for a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
