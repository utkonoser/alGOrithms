package quicksort

type Ordered interface {
	~float64 | ~string | ~int
}

func QuickSort[T Ordered](data []T, low, high int) {
	if low < high {
		var pivot = Partition(data, low, high)
		QuickSort(data, low, pivot)
		QuickSort(data, pivot+1, high)
	}
}

func Partition[T Ordered](data []T, low, high int) int {
	var pivot = data[low]
	var i, j = low, high
	for i < j {
		for data[i] <= pivot && i < high {
			i++
		}
		for data[j] > pivot && j > low {
			j--
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	data[low] = data[j]
	data[j] = pivot
	return j
}

/*
Алгоритм быстрой сортировки является примером алгоритма «разделяй и властвуй».
Мы разделяем исходный срез на два меньших среза и сортируем каждый из них,
продолжая разбивать каждый еще на два, пока в конечном итоге не получим срезы
из двух элементов, которые мы сравниваем друг с другом.
Если исходный срез состоит из n элементов, мы можем выполнить операцию
«разделяй и властвуй» log2n раз (сколько раз мы можем разделить n на 2, чтобы
получить два элемента).
Это предполагает, что мы разделяем срезы, каждый раз разрезая их пополам.
Тщательное изучение статистической суммы показывает, что это не всегда так.
Функция Partition использует свой самый левый элемент в качестве элемента
поворота. Затем он перемещает данные по срезу, чтобы гарантировать, что
элементы слева от опорного элемента меньше, а элементы справа от опорного
элемента больше.
Чем ближе данные к исходной сортировке, тем хуже работает быстрая сортировка.
Полезным фильтром, который можно применить перед быстрой сортировкой, будет
проверка входных данных на предмет того, отсортированы ли они уже. Если это так,
то не выполняйте никакой сортировки.
Поскольку функция Partition равна O(n), алгоритм быстрой сортировки равен O(nlog2n),
когда сортируемые данные еще не отсортированы или близки к сортировке.
*/
