package ejercicios

import (
	"math"
	"untref-ayp2/guia-conjuntos-hashes-diccionarios/dictionary"
	"untref-ayp2/guia-conjuntos-hashes-diccionarios/list"
)

func Traducir(texto string, dict dictionary.Dictionary[string, string]) string {
	return ""
}

func Frecuencia(texto string) *dictionary.Dictionary[string, int] {
	// 🔽 Esta en el test de "hashtable_test"
	hashFn := func(key string) uint {
		const a float64 = 11.0
		l := len(key)
		var hash uint = 0
		for i, c := range key {
			hash += uint(c) * uint(math.Pow(a, float64(l-i-1)))
		}
		return hash
	}
	// 🔼
	dict := dictionary.NewDictionary[string, int](hashFn)
	for c := range texto {
		if char := texto[c]; char != ' ' {
			if dict.Contains(string(char)) {
				a, _ := dict.Get(string(char))
				a++
				dict.Set(string(char), a)
			} else {
				dict.Set(string(char), 1)
			}
		}
	}
	return dict
}

func Interseccion(s1 []string, s2 []string) *list.LinkedList[string] {
	// 🔽 Esta en el test de "hashtable_test"
	hashFn := func(key string) uint {
		const a float64 = 11.0
		l := len(key)
		var hash uint = 0
		for i, c := range key {
			hash += uint(c) * uint(math.Pow(a, float64(l-i-1)))
		}
		return hash
	}
	// 🔼
	dict := dictionary.NewDictionary[string, bool](hashFn)
	for _, v := range s1 {
		dict.Get(v)
	}
	result := list.NewLinkedList[string]()
	for _, v := range s2 {
		if dict.Contains(v) {
			result.Append(v)
		}
	}
	return result
}

func InformacionSolicitada(entrada dictionary.Dictionary[string, []string]) *dictionary.Dictionary[string, []string] {
	//Implementar
	return nil
}
