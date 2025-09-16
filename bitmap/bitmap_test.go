package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitMapNuevoTodoEnCero(t *testing.T) {
	m := NewBitMap()
	assert.Equal(t, uint32(0), m.GetMap())
}

func TestBitMapEncenderUnBit(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	isOn, _ := m.IsOn(1)
	assert.True(t, isOn)
}

func TestBitMapEncenderUnBitEnPosicionNoValida(t *testing.T) {
	m := NewBitMap()
	err := m.On(32)
	assert.EqualError(t, err, "posición no válida")
}

func TestBitMapApagarUnBit(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	m.Off(1)
	isOn, _ := m.IsOn(1)
	assert.False(t, isOn)
}

func TestBitMapApagarUnBitEnPosicionNoValida(t *testing.T) {
	m := NewBitMap()
	err := m.Off(32)
	assert.EqualError(t, err, "posición no válida")
}

func TestBitMapEncenderTodosLosBits(t *testing.T) {
	m := NewBitMap()
	for i := range uint8(32) {
		m.On(i)
	}
	var allOn uint32 = 0xffffffff
	var allOn2 uint32 = 4294967295
	assert.Equal(t, allOn, m.GetMap())
	assert.Equal(t, allOn2, m.GetMap())
}

func TestBitMapApagarTodosLosBits(t *testing.T) {
	m := NewBitMap()
	for i := range uint8(32) {
		m.On(i)
	}
	for i := range uint8(32) {
		m.Off(i)
	}
	assert.Equal(t, uint32(0), m.GetMap())
}

func TestBitMapEstadoDeUnBit(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	isOn, _ := m.IsOn(1)
	assert.True(t, isOn)
}

func TestBitMapEstadoDeUnBitEnPosicionNoValida(t *testing.T) {
	m := NewBitMap()
	_, err := m.IsOn(32)
	assert.EqualError(t, err, "posición no válida")
}

func TestBitMapEncenderVariasVecesUnMismoBitNoLoApaga(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	m.On(1)
	isOn, _ := m.IsOn(1)
	assert.True(t, isOn)
}

func TestBitMapGetMap(t *testing.T) {
	m := NewBitMap()
	m.On(1)
	assert.Equal(t, uint32(2), m.GetMap())
}

func TestBitMapIsOff(t *testing.T) {
	m := NewBitMap()
	m.On(3)
	isOff, _ := m.IsOff(3)
	assert.False(t, isOff) // el bit 3 está encendido

	isOff, _ = m.IsOff(2)
	assert.True(t, isOff) // el bit 2 está apagado
}

func TestBitMapIsOffPosicionNoValida(t *testing.T) {
	m := NewBitMap()
	_, err := m.IsOff(32)
	assert.EqualError(t, err, "posición no válida")
}

func TestBitMapCountOn(t *testing.T) {
	m := NewBitMap()
	assert.Equal(t, 0, m.CountOn()) // ningún bit encendido

	m.On(0)
	m.On(5)
	m.On(31)
	assert.Equal(t, 3, m.CountOn()) // hay 3 bits encendidos
}

func TestBitMapCountOnConTodosLosBits(t *testing.T) {
	m := NewBitMap()
	for i := uint8(0); i < BitmapSize; i++ {
		m.On(i)
	}
	assert.Equal(t, int(BitmapSize), m.CountOn())
}

func TestBitMapStringConBitsEncendidos(t *testing.T) {
	m := NewBitMap()
	m.On(0)  // menos significativo
	m.On(31) // más significativo
	str := m.String()

	assert.Equal(t, "10000000000000000000000000000001", str)
}

func TestBitMapStringTodoApagado(t *testing.T) {
	m := NewBitMap()
	str := m.String()
	assert.Equal(t, "00000000000000000000000000000000", str)
}

func TestBitMapOffVsToggle(t *testing.T) {
	m := NewBitMap()

	m.On(1)
	isOn, _ := m.IsOn(1)
	assert.True(t, isOn)

	m.Off(1)
	isOn, _ = m.IsOn(1)
	assert.False(t, isOn)

	m.Off(1)
	isOn, _ = m.IsOn(1)
	assert.False(t, isOn)

	m.bits ^= 1 << 1
	isOn, _ = m.IsOn(1)
	assert.True(t, isOn)

	m.bits ^= 1 << 1
	isOn, _ = m.IsOn(1)
	assert.False(t, isOn)
}
