package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutAndGetClosed(t *testing.T) {
	ht := NewHashTable[string, int](10, 0.75)

	ht.Put("key1", 10)
	ht.Put("key2", 20)

	value, found := ht.Get("key1")
	assert.True(t, found)
	assert.Equal(t, 10, value)

	value, found = ht.Get("key2")
	assert.True(t, found)
	assert.Equal(t, 20, value)

	ht.Put("key1", 30)
	value, found = ht.Get("key1")
	assert.True(t, found)
	assert.Equal(t, 30, value)
}

func TestCollisionHandlingClosed(t *testing.T) {
	ht := NewHashTable[string, int](5, 0.75)

	ht.Put("a", 1)
	ht.Put("b", 2)
	ht.Put("c", 3)

	val, found := ht.Get("a")
	assert.True(t, found)
	assert.Equal(t, 1, val)

	val, found = ht.Get("b")
	assert.True(t, found)
	assert.Equal(t, 2, val)

	val, found = ht.Get("c")
	assert.True(t, found)
	assert.Equal(t, 3, val)
}

func TestRemoveClosed(t *testing.T) {
	ht := NewHashTable[string, int](10, 0.75)

	ht.Put("key1", 10)
	ht.Put("key2", 20)

	removed := ht.Remove("key1")
	assert.True(t, removed)

	_, found := ht.Get("key1")
	assert.False(t, found)

	// Debe seguir encontrando key2 aunque esté después en la secuencia
	val, found := ht.Get("key2")
	assert.True(t, found)
	assert.Equal(t, 20, val)
}

func TestKeysAndValuesClosed(t *testing.T) {
	ht := NewHashTable[string, int](10, 0.75)

	ht.Put("key1", 10)
	ht.Put("key2", 20)

	keys := ht.Keys()
	assert.ElementsMatch(t, []string{"key1", "key2"}, keys)

	values := ht.Values()
	assert.ElementsMatch(t, []int{10, 20}, values)
}

func TestResizeClosed(t *testing.T) {
	ht := NewHashTable[string, int](2, 0.5)

	ht.Put("key1", 10)
	ht.Put("key2", 20)
	ht.Put("key3", 30) // fuerza resize

	val, found := ht.Get("key1")
	assert.True(t, found)
	assert.Equal(t, 10, val)

	val, found = ht.Get("key2")
	assert.True(t, found)
	assert.Equal(t, 20, val)

	val, found = ht.Get("key3")
	assert.True(t, found)
	assert.Equal(t, 30, val)
}

func TestClearClosed(t *testing.T) {
	ht := NewHashTable[string, int](10, 0.75)

	ht.Put("key1", 10)
	ht.Put("key2", 20)
	ht.Clear()

	assert.Equal(t, uint(0), ht.Size())
	assert.True(t, ht.IsEmpty())
}

func TestIsEmptyAndSizeClosed(t *testing.T) {
	ht := NewHashTable[string, int](10, 0.75)

	assert.True(t, ht.IsEmpty())
	assert.Equal(t, uint(0), ht.Size())

	ht.Put("key1", 10)
	assert.False(t, ht.IsEmpty())
	assert.Equal(t, uint(1), ht.Size())
}
