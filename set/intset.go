package set

type IntSet struct {
	elements map[int]bool
}

func NewIntSet(elements ...int) *IntSet {
	s := &IntSet{elements: make(map[int]bool)}
	for _, v := range elements {
		s.Add(v)
	}
	return s
}

func (s *IntSet) Add(element int) { // O(1)
	s.elements[element] = true
}

func (s *IntSet) Remove(element int) { // O(1)
	delete(s.elements, element)
}

func (s *IntSet) Contains(element int) bool { // O(1)
	return s.elements[element]
}

func (s *IntSet) Size() int { // O(n)
	return len(s.elements)
}

func (s *IntSet) Values() []int {
	values := make([]int, 0, s.Size()) // O(1)
	for k := range s.elements {        // O(n)
		values = append(values, k)
	}
	return values // O(n)
}

// Dado un conjunto A y un conjunto B, la unión de los conjuntos A y B será otro
// conjunto que estará formado por todos los elementos de A, con todos los
// elementos de B sin repetir ningún elemento.
func (s *IntSet) Union(other *IntSet) *IntSet {
	set := NewIntSet(s.Values()...)    // O(1)
	for _, v := range other.Values() { // O(n)
		set.Add(v)
	}
	return set // O(n)
}

// Dado un conjunto A y un conjunto B, la intersección de los conjuntos A y B
// será otro conjunto que estará formado por los elementos de A y los elementos
// de B que sean comunes, los elementos no comunes entre A y B, serán excluidos.
func (s *IntSet) Intersection(other *IntSet) *IntSet {
	set := NewIntSet()             // O(1)
	for _, v := range s.Values() { // O(n)
		if other.Contains(v) { // O(n)
			set.Add(v) // O(n)
		}
	}
	return set // O(n)
}

// Dado un conjunto A y un conjunto B, la diferencia de los conjuntos A y B será
// otro conjunto que estará formado por los elementos de A que no estan
// presentes en B.
func (s *IntSet) Difference(other *IntSet) *IntSet {
	set := NewIntSet()             // O(1)
	for _, v := range s.Values() { // O(n)
		if !other.Contains(v) { // O(n)
			set.Add(v)
		}
	}
	return set // O(n)
}

// Dado un conjunto A y un conjunto B, la diferencia simétrica de los conjuntos
// A y B será otro conjunto que estará formado por todos los elementos no
// comunes a los conjuntos A y B.
func (s *IntSet) SymmetricDifference(other *IntSet) *IntSet {
	set1 := s.Union(other)        // O(n)
	set2 := s.Intersection(other) // O(n)
	return set1.Difference(set2)  // O(n)
}

// Un conjunto A es igual a un conjunto B si ambos conjuntos tienen los mismos
// elementos.
func (s *IntSet) Equal(other *IntSet) bool {
	return s.Size() == other.Size() && s.Subset(other) // O(1)
}

// El conjunto `other` es subconjunto de `s` si todos los elementos de `other`
// están incluidos en `s`.
func (s *IntSet) Subset(other *IntSet) bool {
	for _, v := range other.Values() { // O(n)
		if !s.Contains(v) { // O(n)
			return false
		}
	}
	return true // O(n)
}

// El conjunto `other` es superconjunto de `s` si `other` contiene todos los
// elementos de `s`.
func (s *IntSet) Superset(other *IntSet) bool {
	return other.Subset(s) // O(n)
}
