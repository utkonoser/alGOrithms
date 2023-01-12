package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	numbers := []float64{3.5, -2.4, 12.8, 9.1}
	names := []string{"Zachary", "John", "Jim", "Moe", "Robert"}

	QuickSort[float64](numbers, 0, len(numbers)-1)
	numbersResult := []float64{-2.4, 3.5, 9.1, 12.8}
	if !reflect.DeepEqual(numbers, numbersResult) {
		t.Errorf("%v must be equal %v", numbers, numbersResult)
	}

	QuickSort[string](names, 0, len(names)-1)
	namesResult := []string{"Jim", "John", "Moe", "Robert", "Zachary"}
	if !reflect.DeepEqual(names, namesResult) {
		t.Errorf("%s must be aqual %s", names, namesResult)
	}
}
