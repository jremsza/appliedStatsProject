package main

import (
	"fmt"
	"goStats/appliedStats/data"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"

	"github.com/gonum/stat"
	"github.com/montanaflynn/stats"
)

func statFunc(data [][]float64) []float64 {
	var cor, median1, median2, mean1, mean2 float64

	// Extract the first and second columns
	var col1, col2 []float64
	for _, row := range data {
		col1 = append(col1, row[0])
		col2 = append(col2, row[1])
	}

	// Calculate the statistics
	cor = stat.Correlation(col1, col2, nil)
	median1, _ = stats.Median(col1)
	median2, _ = stats.Median(col2)
	mean1, _ = stats.Mean(col1)
	mean2, _ = stats.Mean(col2)

	return []float64{cor, median1, median2, mean1, mean2}
}

func main() {
	// Start CPU profiling
	f, err := os.Create("profile.prof")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create profile: %v\n", err)
		return
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Fprintf(os.Stderr, "could not start CPU profile: %v\n", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Read the data
	data, err := data.ReadData()
	if err != nil {
		log.Fatalf("error reading data: %v", err)
	}

	// Log data read success
	log.Println("Data read successfully")

	bootstrap := Bootstrap(data, statFunc, 1000)

	// Log bootstrap success
	log.Println("Bootstrap process completed")

	// Calculate the average of the bootstrap results
	average := make([]float64, len(bootstrap[0]))
	for _, statistics := range bootstrap {
		for i, value := range statistics {
			average[i] += value
		}
	}
	for i := range average {
		average[i] /= float64(len(bootstrap))
	}

	statNames := []string{"Correlation", "Sepal Width Median", "Sepal Length Median", "Sepal Width Mean", "Sepal Length Mean"}

	fmt.Println("Bootstrap results:")
	for i, avg := range average {
		fmt.Printf("Statistic %s: %.4f\n", statNames[i], avg)
	}
}

func Bootstrap(data [][]float64, statistic func([][]float64) []float64, nBootstraps int) [][]float64 {
	bootstrappedStatistics := make([][]float64, nBootstraps)
	n := len(data)

	for i := 0; i < nBootstraps; i++ {
		// Generate resampled indices
		indices := make([]int, n)
		for j := range indices {
			indices[j] = rand.Intn(n)
		}

		// Create resampled dataset
		resampledData := make([][]float64, n)
		for j, idx := range indices {
			resampledData[j] = data[idx]
		}

		// Compute statistics on resampled data
		bootstrappedStatistics[i] = statistic(resampledData)
	}

	return bootstrappedStatistics
}
