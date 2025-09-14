// Package set proporciona una implementación de un conjunto genérico basado en un mapa.
// Expone la estructura Set y sus métodos para manipular un conjunto.
package set

import (
	"fmt"
	"strings"
)

// MapSet implementa un conjunto sobre una lista enlazada simple.
type MapSet[T comparable] struct {
	elements map[T]struct{}
}

// NewSet crea un nuevo conjunto vacío y lo inicializa con los elementos especificados.
//
// Uso:
//
//	s1 := set.NewSet[int]() // Crea un nuevo conjunto vacío.
//	s2 := set.NewSet[int](1, 2, 3) // Crea un nuevo conjunto con los elementos 1, 2 y 3.
//
// Parámetros:
//   - `elements`: los elementos con los que inicializar el conjunto.
func NewMapSet[T comparable](elements ...T) *MapSet[T] {
	set := &MapSet[T]{elements: make(map[T]struct{}, len(elements))}
	for _, v := range elements {
		set.Add(v)
	}
	return set
}

// Contains verifica si el conjunto contiene el elemento especificado.
//
// Uso:
//
//	if s.Contains(10) {
//		fmt.Println("El conjunto contiene el elemento 10.")
//	}
//
// Parámetros:
//   - `element`: el elemento a buscar en el conjunto.
//
// Retorna:
//   - `true` si el conjunto contiene el elemento; `false` en caso contrario.
func (s *MapSet[T]) Contains(element T) bool {
	_, exist := s.elements[element]
	return exist
}

// Add agrega los elementos especificados al conjunto.
//
// Uso:
//
//	s.Add(10, 20, 30) // Agrega los elementos 10, 20 y 30 al conjunto.
//
// Parámetros:
//   - `elements`: los elementos a agregar al conjunto.
func (s *MapSet[T]) Add(elements ...T) {
	for _, v := range elements {
		s.elements[v] = struct{}{}
	}
}

// Remove elimina el elemento especificado del conjunto.
//
// Uso:
//
//	s.Remove(10) // Elimina el elemento 10 del conjunto.
//
// Parámetros:
//   - `element`: el elemento a eliminar del conjunto.
func (s *MapSet[T]) Remove(element T) {
	delete(s.elements, element)
}

// Size devuelve la cantidad de elementos en el conjunto.
//
// Uso:
//
//	size := s.Size() // Obtiene la cantidad de elementos en el conjunto.
//
// Retorna:
//   - la cantidad de elementos en el conjunto.
func (s *MapSet[T]) Size() int {
	return len(s.elements)
}

// Values devuelve los elementos del conjunto.
//
// Uso:
//
//	values := s.Values() // Obtiene los elementos del conjunto como un slice.
//
// Retorna:
//   - los elementos del conjunto como un slice.
func (s *MapSet[T]) Values() []T {
	values := make([]T, 0, len(s.elements))
	for k := range s.elements {
		values = append(values, k)
	}
	return values
}

// String devuelve una representación en cadena del conjunto.
//
// Uso:
//
//	fmt.Println(s) // Muestra el conjunto como una cadena.
//
// Retorna:
//   - una representación en cadena del conjunto.
func (s *MapSet[T]) String() string {
	var builder strings.Builder
	builder.WriteString("Set: {")
	var size = s.Size()
	for i, v := range s.Values() {
		builder.WriteString(fmt.Sprintf("%v", v))
		if i+1 < size {
			builder.WriteString(", ")
		}
	}
	builder.WriteRune('}')
	return builder.String()
}
