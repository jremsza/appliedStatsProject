package main

import (
	"fmt"
	"testing"
)

func TestBootstrap(t *testing.T) {
	data := [][]float64{
		{5.1, 3.5, 1.4, 0.2},
		{4.9, 3.0, 1.4, 0.2},
		{4.7, 3.2, 1.3, 0.2},
		{4.6, 3.1, 1.5, 0.2},
		{5.0, 3.6, 1.4, 0.2},
		{5.4, 3.9, 1.7, 0.4},
		{4.6, 3.4, 1.4, 0.3},
		{5.0, 3.4, 1.5, 0.2},
		{4.4, 2.9, 1.4, 0.2},
		{4.9, 3.1, 1.5, 0.1},
		{5.4, 3.7, 1.5, 0.2},
		{4.8, 3.4, 1.6, 0.2},
		{4.8, 3.0, 1.4, 0.1},
		{4.3, 3.0, 1.1, 0.1},
		{5.8, 4.0, 1.2, 0.2},
		{5.7, 4.4, 1.5, 0.4},
		{5.4, 3.9, 1.3, 0.4},
		{5.1, 3.5, 1.4, 0.3},
		{5.7, 3.8, 1.7, 0.3},
		{5.1, 3.8, 1.5, 0.3},
		{5.4, 3.4, 1.7, 0.2},
		{5.1, 3.7, 1.5, 0.4},
		{4.6, 3.6, 1.0, 0.2},
	}
	bootstrap := Bootstrap(data, statFunc, 1000)
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

func BenchmarkBootstrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
