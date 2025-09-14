package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMapSet(t *testing.T) {
	set := NewMapSet[int]()

	assert.NotNil(t, set)
	assert.Equal(t, 0, set.Size())
}

func TestMapSetAdd(t *testing.T) {
	set := NewMapSet[int]()

	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestMapSetAddMultiple(t *testing.T) {
	set := NewMapSet[int]()

	set.Add(1, 2, 3)
	assert.Equal(t, 3, set.Size())
}

func TestMapSetAddExistenteNoRepite(t *testing.T) {
	set := NewMapSet[int]()

	set.Add(1)
	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestMapSetContains(t *testing.T) {
	set := NewMapSet[int]()
	set.Add(1)

	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
}

func TestMapSetRemove(t *testing.T) {
	set := NewMapSet[int]()
	set.Add(1)
	set.Add(2)
	assert.True(t, set.Contains(1))
	assert.Equal(t, 2, set.Size())

	set.Remove(1)
	assert.False(t, set.Contains(1))
	assert.Equal(t, 1, set.Size())
}

func TestMapSetRemoveNonExistent(t *testing.T) {
	set := NewMapSet[int]()
	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Remove(2)
	assert.Equal(t, 1, set.Size())
}

func TestMapSetSize(t *testing.T) {
	set := NewMapSet[int]()
	assert.Equal(t, 0, set.Size())

	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Add(2)
	assert.Equal(t, 2, set.Size())
}

func TestMapSetValuesOnAnEmptySet(t *testing.T) {
	set := NewMapSet[int]()
	values := set.Values()

	assert.Empty(t, values)
}

func TestMapSetValuesOnANonEmptySet(t *testing.T) {
	set := NewMapSet(1, 2)
	values := set.Values()

	assert.Len(t, values, 2)
	assert.ElementsMatch(t, []int{1, 2}, values)
}

func TestMapSetStringEnSetVacio(t *testing.T) {
	set := NewMapSet[int]()
	assert.Equal(t, "Set: {}", set.String())
}

func TestMapSetStringEnSetNoVacio(t *testing.T) {
	set := NewMapSet(1, 2)
	possibleRepresentations := []string{
		"Set: {1, 2}",
		"Set: {2, 1}",
	}

	assert.Contains(t, possibleRepresentations, set.String())
}
