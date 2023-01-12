package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	numbers := []float64{3.5, -2.4, 12.8, 9.1}
	names := []string{"Zachary", "John", "Moe", "Jim", "Robert"}
	BubbleSort[float64](numbers)
	numbersResult := []float64{-2.4, 3.5, 9.1, 12.8}
	if !reflect.DeepEqual(numbers, numbersResult) {
		t.Errorf("%v must be equal %v", numbers, numbersResult)
	}
	BubbleSort[string](names)
	namesResult := []string{"Jim", "John", "Moe", "Robert", "Zachary"}
	if !reflect.DeepEqual(names, namesResult) {
		t.Errorf("%s must be equal %s", names, namesResult)
	}
}
