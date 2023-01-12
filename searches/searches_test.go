package searches

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLinearSearch(t *testing.T) {
	data := make([]float64, Size)
	for i := 0; i < Size; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	start := time.Now()
	result := linearSearch[float64](data, 54.0)
	elapsed := time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using linear search: ", elapsed)
	fmt.Println("Result of search is: ", result)

	start = time.Now()
	result = linearSearch[float64](data, data[Size/2])
	elapsed = time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using linear search: ", elapsed)
	fmt.Println("Result of search is: ", result)
}

func TestConcurrentSearch(t *testing.T) {
	data := make([]float64, Size)
	for i := 0; i < Size; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	start := time.Now()
	result := concurrentSearch[float64](data, 54.0)
	elapsed := time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using concurrent search: ", elapsed)
	fmt.Println("Result of search is: ", result)

	start = time.Now()
	result = concurrentSearch[float64](data, data[Size/2])
	elapsed = time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using concurrent search: ", elapsed)
	fmt.Println("Result of search is: ", result)
}

func TestBinarySearch(t *testing.T) {
	data := make([]float64, Size)
	for i := 0; i < Size; i++ {
		data[i] = float64(i)
	}
	start := time.Now()
	result := binarySearch[float64](data, -1.0)
	elapsed := time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using binary search: ", elapsed)
	fmt.Println("Result of search is: ", result)

	start = time.Now()
	result = binarySearch[float64](data, data[Size/2])
	elapsed = time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using binary search: ", elapsed)
	fmt.Println("Result of search is: ", result)
}
