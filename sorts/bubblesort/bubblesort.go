package bubblesort

// Generic bubble sort

type Ordered interface {
	~float64 | ~int | ~string
}

func BubbleSort[T Ordered](data []T) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

/*
Тип Ordered может включать гораздо больше базовых типов.
Символ тильды перед каждым из основных типов означает, что любой
пользовательский тип, использующий данный базовый тип, считается
Ordered. Элементы сравниваются последовательно и меняются местами,
если условие не выполняется. На каждой итерации внешнего цикла
наибольшее значение «перескакивает» в крайнее правое положение в срезе.
Эта позиция не учитывается при следующей итерации внутреннего цикла
из-за `j < n – 1 – i`. Вложенные циклы for, делают этот алгоритм O(n^2).
В общем, k вложенных циклов производят алгоритм O (n^k).
Пузырьковая сортировка наиболее эффективна, когда сортируемый срез
уже отсортирован, и медленнее всего, когда срез находится в обратном порядке.
*/