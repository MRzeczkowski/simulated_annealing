package main

import (
	"fmt"
	"math"
	"math/rand"
)

func rastrigin(x []float64) float64 {
	n := len(x)
	sum := 10.0 * float64(n)
	for _, xi := range x {
		sum += (xi*xi - 10.0*math.Cos(2*math.Pi*xi))
	}
	return sum
}

var Min = -5.12
var Max = 5.12

func simulatedAnnealing(dimensions int, maxIterations int, initialTemp float64, minTemp float64, coolingRate float64, maxCount int, coolingScheme string) []float64 {

	currentSolution := make([]float64, dimensions)
	for i := range currentSolution {
		currentSolution[i] = rand.Float64()*(Max-Min) + Min
	}

	currentEnergy := rastrigin(currentSolution)
	bestSolution := make([]float64, dimensions)
	copy(bestSolution, currentSolution)
	bestEnergy := currentEnergy

	temperature := initialTemp
	count := 0

	for temperature > minTemp && count < maxIterations {
		for i := 0; i < maxCount; i++ {
			newSolution := make([]float64, dimensions)
			for j := range newSolution {
				newSolution[j] = currentSolution[j] + (rand.Float64()*2 - 1)
			}

			newEnergy := rastrigin(newSolution)

			if newEnergy < bestEnergy {
				copy(bestSolution, newSolution)
				bestEnergy = newEnergy
			}

			if newEnergy < currentEnergy || rand.Float64() < math.Exp((currentEnergy-newEnergy)/temperature) {
				copy(currentSolution, newSolution)
				currentEnergy = newEnergy
			}

			if newEnergy < bestEnergy {
				copy(bestSolution, newSolution)
				bestEnergy = newEnergy
			}

			count++
		}

		switch coolingScheme {
		case "geometric":
			temperature *= coolingRate
		case "linear":
			temperature -= coolingRate
		case "exponential":
			temperature = initialTemp * math.Pow(minTemp/initialTemp, float64(count)/float64(maxIterations))
		case "logarithmic":
			temperature = initialTemp / math.Log(float64(count)+1)
		case "harmonic":
			A := (initialTemp - minTemp) * float64(maxIterations+1) / float64(maxIterations)
			temperature = A/float64(count+1) + initialTemp - A
		default:
			fmt.Println("Unknown cooling scheme, using geometric by default.")
			temperature *= coolingRate
		}
	}

	return bestSolution
}

func runExperiment(dimensions int, maxIterations int, initialTemp float64, minTemp float64, coolingRate float64, maxCount int, coolingScheme string) float64 {
	solution := simulatedAnnealing(dimensions, maxIterations, initialTemp, minTemp, coolingRate, maxCount, coolingScheme)
	return rastrigin(solution)
}

func main() {
	dimensions := 3
	maxIterations := 10000
	numberOfTests := 10

	fmt.Println("Simulation Parameters:")
	fmt.Println("- Number of dimensions:", dimensions)
	fmt.Println("- Max iterations per test:", maxIterations)
	fmt.Println("- Number of tests per cooling scheme:", numberOfTests)
	fmt.Println()

	coolingSchemes := []string{"geometric", "linear", "exponential", "logarithmic", "harmonic"}

	bestResults := make(map[string]float64)
	bestParams := make(map[string]struct {
		initialTemp float64
		minTemp     float64
		coolingRate float64
		maxCount    int
	})

	for _, scheme := range coolingSchemes {
		bestAverageResult := math.Inf(1)
		var bestParam struct {
			initialTemp float64
			minTemp     float64
			coolingRate float64
			maxCount    int
		}

		// fmt.Println("| Cooling scheme | Initial temperature | Min temperature | Cooling rate | Max iterations at temperature level | Average result |")
		// fmt.Println("|-|-|-|-|-|-|")

		for initTemp := 500.0; initTemp <= 1500.0; initTemp += 100.0 {
			for minTemp := 0.1; minTemp <= 1.0; minTemp += 0.1 {
				for coolRate := 0.8; coolRate <= 0.99; coolRate += 0.05 {
					for maxCount := 25; maxCount <= 200; maxCount += 25 {
						// Skip cooling rate changes for schemes where it is not applicable
						if (scheme == "exponential" || scheme == "logarithmic" || scheme == "harmonic") && coolRate != 0.8 {
							continue
						}

						sumResults := 0.0

						for i := 0; i < numberOfTests; i++ {
							sumResults += runExperiment(dimensions, maxIterations, initTemp, minTemp, coolRate, maxCount, scheme)
						}

						averageResult := sumResults / float64(numberOfTests)
						// fmt.Printf("| %s | %.2f | %.2f | %.2f | %d | %.4f |\n", scheme, initTemp, minTemp, coolRate, maxCount, result)

						if averageResult < bestAverageResult {
							bestAverageResult = averageResult
							bestParam = struct {
								initialTemp float64
								minTemp     float64
								coolingRate float64
								maxCount    int
							}{initTemp, minTemp, coolRate, maxCount}
						}
					}
				}
			}
		}

		bestResults[scheme] = bestAverageResult
		bestParams[scheme] = bestParam
	}

	fmt.Println("Best parameters for each cooling scheme")
	fmt.Println("| Cooling scheme | Initial temperature | Min temperature | Cooling rate | Max iterations at temperature level | Average result |")
	fmt.Println("|-|-|-|-|-|-|")
	for scheme, result := range bestResults {
		param := bestParams[scheme]
		if scheme == "exponential" || scheme == "logarithmic" || scheme == "harmonic" {
			fmt.Printf("| %s | %.2f | %.2f | None | %d | %.4f |\n", scheme, param.initialTemp, param.minTemp, param.maxCount, result)
		} else {
			fmt.Printf("| %s | %.2f | %.2f | %.2f | %d | %.4f |\n", scheme, param.initialTemp, param.minTemp, param.coolingRate, param.maxCount, result)
		}
	}
}
