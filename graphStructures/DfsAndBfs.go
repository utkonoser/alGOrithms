package graphStructures

// OrderedStringer - сущности этого типа должны быть comparable и
// иметь строковое представление с помощью метода String().
type OrderedStringer interface {
	comparable
	String() string
}

// Vertex - вершина графа содержит значение Key и мапу Neighbours, отображающую ключи
// на указатели на другие вершины.
type Vertex[T OrderedStringer] struct {
	Key       T
	Neighbors map[T]*Vertex[T]
}

// Graph - структура графа
type Graph[T OrderedStringer] struct {
	Vertices map[T]*Vertex[T]
}

// visitation - слайс, представляющий обход графа
var visitation []string

// NewVertex - создает вершину графа
func NewVertex[T OrderedStringer](key T) *Vertex[T] {
	return &Vertex[T]{
		Key:       key,
		Neighbors: map[T]*Vertex[T]{},
	}
}

// NewGraph - создает пустой новый граф
func NewGraph[T OrderedStringer]() *Graph[T] {
	return &Graph[T]{Vertices: map[T]*Vertex[T]{}}
}

// AddVertex - создает новую вершину в графе g
func (g *Graph[T]) AddVertex(key T) {
	vertex := NewVertex(key)
	g.Vertices[key] = vertex
}

// AddEdge - берет два ключа, присваивает их полю Vertices
// графа и присваивает Neighbors вершинам.
func (g *Graph[T]) AddEdge(key1, key2 T) {
	vertex1 := g.Vertices[key1]
	vertex2 := g.Vertices[key2]
	if vertex1 == nil || vertex2 == nil {
		return
	}
	vertex1.Neighbors[vertex2.Key] = vertex2
	g.Vertices[vertex1.Key] = vertex1
	g.Vertices[vertex2.Key] = vertex2
}

func (g *Graph[T]) String() string {
	result := ""
	for i := 0; i < len(visitation); i++ {
		result += " " + visitation[i]
	}
	return result
}

type String string

func (str String) String() string {
	return string(str)
}

// DepthFirstSearch - это поиск в глубину. Этот метод перемещается наружу
// от начальной вершины и далее идет непосредственно к самой дальней
// еще не посещенной вершине.
func (g *Graph[T]) DepthFirstSearch(start *Vertex[T], visited map[T]bool) {
	if start == nil {
		return
	}
	visited[start.Key] = true
	visitation = append(visitation, start.Key.String())

	for _, v := range start.Neighbors {
		if visited[v.Key] {
			continue
		}
		g.DepthFirstSearch(v, visited)
	}
}

// Queue - структура очереди
type Queue[T any] struct {
	items []T
}

func (queue *Queue[T]) Insert(item T) {
	queue.items = append(queue.items, item)
}

func (queue *Queue[T]) Remove() T {
	returnValue := queue.items[0]
	queue.items = queue.items[1:]
	return returnValue
}

// BreadthFirstSearch - поиск в ширину. Этот метод посещает
// вершины, близкие к начальной, медленно двигаясь наружу.
// Здесь не используется рекурсия. Очередь используется для хранения
// вершин, которые являются соседними посещенными вершинами. Эти соседние вершины
// проходятся первыми. Итак, как следует из названия этого метода, этот обход
// движется к соседним вершинам, медленно удаляясь от начальной вершины.
func (g *Graph[T]) BreadthFirstSearch(start *Vertex[T], visited map[T]bool) {
	if start == nil {
		return
	}
	queue := Queue[*Vertex[T]]{}

	current := start
	for {
		if !visited[current.Key] {
			visitation = append(visitation, current.Key.String())
		}
		visited[current.Key] = true

		for _, v := range current.Neighbors {
			if !visited[v.Key] {
				queue.Insert(v)
			}
		}
		if len(queue.items) > 0 {
			current = queue.Remove()
		} else {
			break
		}
	}
}
