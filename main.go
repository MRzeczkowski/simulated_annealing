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

func runExperiment(dimensions int, maxIterations int, initialTemp float64, minTemp float64, coolingRate float64, maxCount int, coolingScheme string) {
	fmt.Printf("Running experiment with cooling scheme: %s, initialTemp: %.2f, minTemp: %.2f, coolingRate: %.3f, maxCount: %d\n", coolingScheme, initialTemp, minTemp, coolingRate, maxCount)
	solution := simulatedAnnealing(dimensions, maxIterations, initialTemp, minTemp, coolingRate, maxCount, coolingScheme)
	fmt.Printf("Found solution: %v\n", solution)
	fmt.Printf("Rastrigin value: %f\n\n", rastrigin(solution))
}

func main() {
	dimensions := 3
	maxIterations := 1000

	experiments := []struct {
		coolingScheme string
		initialTemp   float64
		minTemp       float64
		coolingRate   float64
		maxCount      int
	}{
		{"geometric", 1000.0, 0.1, 0.9, 100},
		{"linear", 1000.0, 0.1, 0.5, 100},
		{"exponential", 1000.0, 0.1, 0.0, 100},
		{"logarithmic", 1000.0, 0.1, 0.0, 100},
		{"harmonic", 1000.0, 0.1, 0.0, 100},
	}

	for _, exp := range experiments {
		runExperiment(dimensions, maxIterations, exp.initialTemp, exp.minTemp, exp.coolingRate, exp.maxCount, exp.coolingScheme)
	}
}
