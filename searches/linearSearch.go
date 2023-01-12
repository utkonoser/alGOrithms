package searches

// Самый простой поисковый алгоритм для поиска среза — это линейный поиск. Мы
// последовательно перебираем все элементы среза, пока не найдем совпадение или
// не завершим поиск всех элементов среза. Сложность O(n)

const Size = 100_000_000

type Ordered interface {
	~float64 | ~int | ~string
}

func linearSearch[T Ordered](slice []T, target T) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return true
		}
	}
	return false
}
