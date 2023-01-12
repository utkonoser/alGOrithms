package searches

// Если данные в искомом срезе отсортированы, можно использовать алгоритм
// бинарного поиска. Этот алгоритм имеет сложность O(log2n), так как
// пространство поиска уменьшается вдвое во время каждой итерации.

func binarySearch[T Ordered](slice []T, target T) bool {
	low := 0
	high := len(slice) - 1

	for low <= high {
		mid := (low + high) / 2

		if slice[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low == len(slice) || slice[low] != target {
		return false
	}
	return true
}
