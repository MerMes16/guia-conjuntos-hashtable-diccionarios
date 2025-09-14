package ejercicios

import "untref-ayp2/guia-conjuntos-hashes-diccionarios/set"

func EliminarRepetidos[T comparable](array []T) []T {
	set := set.NewListSet[T]()
	for _, v := range array {
		set.Add(v)
	}
	return set.Values()
}

func InterseccionMultiple[T comparable](sets ...*set.ListSet[T]) *set.ListSet[T] {
	// Implementar
	return nil
}
