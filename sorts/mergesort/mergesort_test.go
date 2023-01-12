package mergesort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	start := time.Now()
	result := MergeSort[float64](data)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time for Mergesort = ", elapsed)
	if IsSorted(result) {
		fmt.Println("Is sorted: ", IsSorted(result))
	} else {
		t.Error("fail sorted")
	}

}
