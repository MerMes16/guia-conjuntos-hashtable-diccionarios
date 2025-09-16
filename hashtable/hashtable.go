// hashtable proporciona una implementación de una tabla hash cerrada cuyas
// claves son cadenas de caracteres y los valores pueden ser de cualquier tipo.
// La tabla utiliza un arreglo para almacenar pares clave-valor.
package hashtable

import (
	"fmt"
	"hash/maphash"
)

// hashTableEntry representa una entrada en la tabla hash, que contiene una
// clave y su valor asociado.
type hashTableEntry[K comparable, V any] struct {
	key   K
	value V
}

// HashTable es una tabla hash cerrada que utiliza un arreglo para almacenar
// elementos. La tabla solo soporta string como claves y cualquier tipo como
// valores. En cada posición del arreglo se almacena un par clave-valor.
type HashTable[K comparable, V any] struct {
	// arreglo de entradas de la tabla hash.
	buckets []*hashTableEntry[K, V]
	// size es el número de elementos en la tabla.
	size uint
	// capacity es la capacidad de la tabla.
	capacity uint
	// loadFactor es el factor de carga de la tabla.
	loadFactor float32
	// threshold es el umbral de carga para redimensionar la tabla.
	threshold uint
	// semilla crea una semilla aleatoria para el maphash
	semilla maphash.Seed
}

// NewHashTable crea una nueva tabla de hash cerrada con la capacidad y el
// factor de carga especificados.
//
// - Si la capacidad es igual a 0, se establece en 17.
//
// - Si el factor de carga es menor o igual a 0 o mayor que 1, se establece en
// 0.75.
//
// - Si la capacidad no es un número primo, se redimensiona a la siguiente
// capacidad primo mayor o igual a la capacidad especificada.
func NewHashTable[K comparable, V any](capacity uint, loadFactor float32) *HashTable[K, V] {
	if capacity == 0 {
		capacity = 17
	}
	if loadFactor <= 0 || loadFactor > 1 {
		loadFactor = 0.75
	}
	if !isPrime(capacity) {
		capacity = nextPrime(capacity)
	}
	return &HashTable[K, V]{
		buckets:    make([]*hashTableEntry[K, V], capacity),
		size:       0,
		capacity:   capacity,
		loadFactor: loadFactor,
		threshold:  uint(float32(capacity) * loadFactor),
		semilla:    maphash.MakeSeed(),
	}
}

// Put agrega un nuevo par clave-valor a la tabla de hash. Si la clave ya
// existe, actualiza el valor asociado a la clave.
//
// Devuelve true si se agregó o actualizó el elemento, false si la clave es nula.
//
// - Si la tabla de hash está llena, se redimensiona automáticamente.
//
// - Si la clave es nula, no se agrega nada.
func (ht *HashTable[K, V]) Put(key K, value V) bool {
	// Si la clave es nula, no se agrega nada.
	var zeroValue K
	if key == zeroValue {
		return false
	}
	// Si la tabla de hash está llena, redimensionamos.
	if ht.size >= ht.threshold {
		ht.resize()
	}

	index := ht.hasheada(key) % ht.capacity
	for {
		if ht.buckets[index] == nil || ht.buckets[index].key == zeroValue {
			// Si el bucket está vacío o la clave es nula, insertamos el nuevo elemento.
			ht.buckets[index] = &hashTableEntry[K, V]{key: key, value: value}
			ht.size++
			return true
		} else if ht.buckets[index].key == key {
			// Si la clave ya existe, actualizamos el valor.
			ht.buckets[index].value = value
			return true
		}
		// Si el bucket está ocupado y la clave no coincide, probamos el siguiente índice.
		index = (index + 1) % ht.capacity
	}
}

// Get devuelve el valor asociado a la clave dada y true para indicar que
// encontró la clave buscada.
//
// - Si la clave no existe o es nula, devuelve false y un valor nulo.
func (ht *HashTable[K, V]) Get(key K) (V, bool) {
	index, exists := ht.getIndex(key)
	if !exists {
		var zeroValue V
		return zeroValue, exists
	}
	return ht.buckets[index].value, exists
}

// Remove elimina el par clave-valor asociado a la clave dada.
//
// Devuelve true si se eliminó el elemento, false si la clave no existe.
func (ht *HashTable[K, V]) Remove(key K) bool {
	index, exists := ht.getIndex(key)
	if exists {
		var zeroKValue K
		ht.buckets[index].key = zeroKValue
		var zeroVValue V
		ht.buckets[index].value = zeroVValue
		ht.size--
	}
	return exists
}

// Keys devuelve una lista de todas las claves en la tabla de hash.
func (ht *HashTable[K, V]) Keys() []K {
	keys := make([]K, 0, ht.size)
	var zeroKValue K
	for _, node := range ht.buckets {
		if node != nil && node.key != zeroKValue {
			keys = append(keys, node.key)
		}
	}
	return keys
}

// Values devuelve una lista de todos los valores en la tabla de hash.
func (ht *HashTable[K, V]) Values() []V {
	values := make([]V, 0, ht.size)
	var zeroKValue K
	for _, node := range ht.buckets {
		if node != nil && node.key != zeroKValue {
			values = append(values, node.value)
		}
	}
	return values
}

// Size devuelve el número de elementos en la tabla de hash.
func (ht *HashTable[K, V]) Size() uint {
	return ht.size
}

// IsEmpty devuelve true si la tabla de hash está vacía, false en caso contrario.
func (ht *HashTable[K, V]) IsEmpty() bool {
	return ht.size == 0
}

// Clear elimina todos los elementos de la tabla de hash.
func (ht *HashTable[K, V]) Clear() {
	ht.buckets = make([]*hashTableEntry[K, V], ht.capacity)
	ht.size = 0
}

// String devuelve una representación en cadena de la tabla de hash.
func (ht *HashTable[K, V]) String() string {
	result := "{"
	var zeroKValue K
	for _, node := range ht.buckets {
		if node != nil && node.key != zeroKValue {
			result += fmt.Sprintf("%v: %v", node.key, node.value) + ", "
		}
	}
	if len(result) > 1 {
		result = result[:len(result)-2]
	}
	result += "}"
	return result
}

func (ht *HashTable[K, V]) hasheada(key K) uint {
	var h maphash.Hash
	h.SetSeed(ht.semilla)
	_, _ = h.WriteString(fmt.Sprintf("%v", key))
	return uint(h.Sum64())
}

// Funciones privadas //////////////////////////////////////////////////////////

// getIndex devuelve el índice del bucket para una clave dada y un booleano que
// indica si la clave existe.
func (ht *HashTable[K, V]) getIndex(key K) (uint, bool) {
	var zeroKValue K
	if key == zeroKValue {
		return 0, false
	}
	for index := ht.hasheada(key) % ht.capacity; ht.buckets[index] != nil; index = (index + 1) % ht.capacity {
		if ht.buckets[index].key == key {
			return index, true
		}
	}
	return 0, false
}

// resize redimensiona la tabla de hash y reubica todos los elementos en la
// nueva tabla.
//
// El nuevo tamaño es el siguiente número primo mayor o igual al doble de la
// capacidad actual.
func (ht *HashTable[K, V]) resize() {
	newCapacity := nextPrime(ht.capacity * 2)
	newBuckets := make([]*hashTableEntry[K, V], newCapacity)
	var zeroKValue K
	// Reinsertar todos los elementos en el nuevo arreglo, manejando colisiones
	for _, node := range ht.buckets {
		if node != nil && node.key != zeroKValue {
			index := ht.hasheada(node.key) % newCapacity
			for newBuckets[index] != nil {
				// Resolver colisiones con prueba lineal
				index = (index + 1) % newCapacity
			}
			newBuckets[index] = node
		}
	}

	// Actualizar los atributos de la tabla hash
	ht.buckets = newBuckets
	ht.capacity = newCapacity
	ht.threshold = uint(float32(newCapacity) * ht.loadFactor)
}

// nextPrime devuelve el siguiente número primo mayor o igual a n.
func nextPrime(n uint) uint {
	if n <= 1 {
		return 2
	}
	for i := n; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}

// isPrime devuelve true si n es un número primo, false en caso contrario.
func isPrime(n uint) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := uint(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
