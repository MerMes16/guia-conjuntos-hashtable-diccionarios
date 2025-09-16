package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTraducir(t *testing.T) {
	// dic := dictionary.NewDictionary[string, string]()
	// dic.Put("Dungeons", "Calabozos")
	// dic.Put("Dragons", "Dragones")

	// salida := Traducir("Dungeons", *dic)
	// assert.Equal(t, "Calabozos", salida)

	// salida = Traducir("Dwarf", *dic)
	// assert.Equal(t, "error", salida)

	// salida = Traducir("Dungeons & Dragons", *dic)
	// assert.Equal(t, "Calabozos error Dragones", salida)
}

func TestFrecuencia(t *testing.T) {
	dict := Frecuencia("ahora")
	require.NotNil(t, dict)
	assert.Equal(t, 4, dict.Size())
	val, err := dict.Get("a")
	require.NoError(t, err)
	assert.Equal(t, 2, val)

	val, err = dict.Get("h")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = dict.Get("o")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = dict.Get("r")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	dict = Frecuencia("hoy es lunes")
	assert.Equal(t, 8, dict.Size())
	val, err = dict.Get("h")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = dict.Get("o")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = dict.Get("y")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = dict.Get("e")
	require.NoError(t, err)
	assert.Equal(t, 2, val)

	val, err = dict.Get("s")
	require.NoError(t, err)
	assert.Equal(t, 2, val)

	val, err = dict.Get("l")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = dict.Get("u")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = dict.Get("n")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

}

func TestInterseccion(t *testing.T) {
	s1 := []string{"A", "B", "C"}
	s2 := []string{"A", "D", "E"}
	list := Interseccion(s1, s2)
	require.NotNil(t, list)
	assert.Equal(t, 1, list.Size())
	assert.NotNil(t, list.Find("A"))

	s1 = []string{"A", "B"}
	s2 = []string{"C", "D"}
	list = Interseccion(s1, s2)
	assert.Equal(t, 0, list.Size())

	s1 = []string{"A"}
	s2 = make([]string, 0)
	list = Interseccion(s1, s2)
	assert.Equal(t, 0, list.Size())

	s1 = make([]string, 0)
	s2 = []string{"A", "B"}
	list = Interseccion(s1, s2)
	assert.Equal(t, 0, list.Size())
}

// func TestInformacionSolicitada(t *testing.T) {
// 	entrada := dictionary.NewDictionary[string, []string]()
// 	sl1 := []string{"Ana", "Pedro"}
// 	sl2 := []string{"Ana"}
// 	entrada.Put("Mie 10", sl1)
// 	entrada.Put("Vie 12", sl2)
// 	salida := InformacionSolicitada(*entrada)
// 	require.NotNil(t, salida)
// 	valPedro, _ := salida.Get("Pedro")
// 	assert.ElementsMatch(t, []string{"Mie 10"}, valPedro)
// 	valAna, _ := salida.Get("Ana")
// 	assert.ElementsMatch(t, []string{"Mie 10", "Vie 12"}, valAna)
// }
