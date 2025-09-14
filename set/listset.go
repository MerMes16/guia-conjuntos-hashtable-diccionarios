// Package set proporciona una implementación de un conjunto genérico basado en una lista dinámica.
// Expone la estructura ListSet y sus métodos para manipular un conjunto.
package set

import (
	"fmt"
	"strings"
	"untref-ayp2/guia-conjuntos-hashes-diccionarios/list"
)

// ListSet implementa un conjunto sobre una lista enlazada simple.
type ListSet[T comparable] struct {
	elements *list.LinkedList[T]
}

// NewListSet crea un nuevo conjunto vacío y lo inicializa con los elementos especificados.
//
// Uso:
//
//	s1 := set.NewListSet[int]() // Crea un nuevo conjunto vacío.
//	s2 := set.NewListSet[int](1, 2, 3) // Crea un nuevo conjunto con los elementos 1, 2 y 3.
//
// Parámetros:
//   - `elements`: los elementos con los que inicializar el conjunto.
func NewListSet[T comparable](elements ...T) *ListSet[T] {
	set := &ListSet[T]{elements: list.NewLinkedList[T]()}
	for _, v := range elements {
		if !set.Contains(v) {
			set.elements.Append(v)
		}
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
func (s *ListSet[T]) Contains(element T) bool {
	for current := s.elements.Head(); current != nil; current = current.Next() {
		if current.Data() == element {
			return true
		}
	}
	return false
}

// Add agrega los elementos especificados al conjunto.
//
// Uso:
//
//	s.Add(10, 20, 30) // Agrega los elementos 10, 20 y 30 al conjunto.
//
// Parámetros:
//   - `elements`: los elementos a agregar al conjunto.
func (s *ListSet[T]) Add(elements ...T) {
	for _, v := range elements {
		if !s.Contains(v) {
			s.elements.Append(v)
		}
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
func (s *ListSet[T]) Remove(element T) {
	for current := s.elements.Head(); current != nil; current = current.Next() {
		if current.Data() == element {
			s.elements.Remove(current.Data())
			return
		}
	}
}

// Size devuelve la cantidad de elementos en el conjunto.
//
// Uso:
//
//	size := s.Size() // Obtiene la cantidad de elementos en el conjunto.
//
// Retorna:
//   - la cantidad de elementos en el conjunto.
func (s *ListSet[T]) Size() int {
	return s.elements.Size()
}

// Values devuelve los elementos del conjunto.
//
// Uso:
//
//	values := s.Values() // Obtiene los elementos del conjunto como un slice.
//
// Retorna:
//   - los elementos del conjunto como un slice.
func (s *ListSet[T]) Values() []T {
	values := make([]T, 0, s.Size())
	for current := s.elements.Head(); current != nil; current = current.Next() {
		values = append(values, current.Data())
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
func (s *ListSet[T]) String() string {
	var builder strings.Builder
	builder.WriteString("Set: {")
	var size = s.Size()
	for i, v := range s.Values() {
		builder.WriteString(fmt.Sprintf("%v", v))
		if i+1 < size {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("}")
	return builder.String()
}
