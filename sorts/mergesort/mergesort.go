package mergesort

/*
Следующий алгоритм сортировки, который мы исследуем — это классический
алгоритм сортировки слиянием. Это алгоритм «разделяй и властвуй», как
и быстрая сортировка. Мы заменяем исходный срез двумя слайсами, каждый
размером в половину исходного размера среза. Каждый из этих полуфрагментов
далее делится на четвертинки, и этот шаблон продолжается до тех пор, пока мы
не получим фрагменты размера 1. Используя рекурсию, мы сплетаем фрагменты вместе,
объединяя их.
*/

const size = 50_000_000

type Ordered interface {
	~float64 | ~int | ~string
}

func IsSorted[T Ordered](data []T) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func InsertSort[T Ordered](data []T) {
	i := 1
	for i < len(data) {
		h := data[i]
		j := i - 1
		for j >= 0 && h < data[j] {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = h
		i++
	}
}

func Merge[T Ordered](left, right []T) []T {
	result := make([]T, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
	return result
}

func MergeSort[T Ordered](data []T) []T {
	if len(data) > 100 {
		middle := len(data) / 2
		left := data[:middle]
		right := data[middle:]
		data = Merge(MergeSort(right), MergeSort(left))
	} else {
		InsertSort(data)
	}
	return data
}

/*
Этот алгоритм проще для понимания, чем быстрая сортировка. Функция Merge создает новый срез,
результат, путем слияния элементов из двух входных срезов, левого и правого, так что
результат сортируется.
Рекурсивная функция MergeSort разбивает входной массив на левый и правый и вызывает Merge
по результатам рекурсивного вызова MergeSort.
Следует отметить, что MergeSort не выполняет сортировку на месте, как это делает быстрая
сортировка. Это требует дополнительного выделения памяти по сравнению с быстрой сортировкой.
Поскольку Merge равно O(n) и в MergeSort имеется log2n рекурсивных вызовов, сложность
MergeSort составляет O(nlog2n).
*/
