package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var workingCars float64
	workingCars = float64(productionRate) * (float64(successRate) / 100)
	return workingCars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var workingCars float64
	workingCars = float64(productionRate) * (float64(successRate) / 6000)
	return int(workingCars)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	return uint((carsCount/10)*95000 + (carsCount%10)*10000)
}
