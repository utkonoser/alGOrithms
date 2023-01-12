package searches

// Конкурентная вариация линейного поиска. Сложность O(n)

import "runtime"

func searchSegment[T Ordered](slice []T, target T, a, b int, ch chan<- bool) {
	for i := a; i < b; i++ {
		if slice[i] == target {
			ch <- true
		}
	}
	ch <- false
}

func concurrentSearch[T Ordered](data []T, target T) bool {
	ch := make(chan bool)
	numSegments := runtime.NumCPU()
	segmentSize := int(float64(len(data)) / float64(numSegments))
	for index := 0; index < numSegments; index++ {
		go searchSegment(data, target, index*segmentSize, index*segmentSize+segmentSize, ch)
	}
	num := 0
	for {
		select {
		case value := <-ch:
			if value == true {
				return true
			}
			num++
			if num == numSegments {
				return false
			}
		}
	}
	return false
}
