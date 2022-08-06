package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCars := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(workingCars) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	oneCarPrice := 10000
	tenCarsPrice := 95000
	cost := (carsCount / 10 * tenCarsPrice) + (carsCount % 10 * oneCarPrice)
	return uint(cost)
}
