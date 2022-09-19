package cola_test

import (
	TDACola "cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	t.Log("Hagamos la prueba de cola vacía...")
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia(), "La cola vacía deber ser true")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestUnElemento(t *testing.T) {
	t.Log("Probemos que funcione para un solo elemento")
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(5)
	require.Equal(t, 5, cola.VerPrimero(), "El primer elemento de la cola debe ser 5")
	cola.Desencolar()
}

func TestOtroTipo(t *testing.T) {
	t.Log("Hagamos la prueba ahora para strings...")
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("Hola")
	cola.Encolar("Mundo!")
	require.Equal(t, "Hola", cola.Desencolar(), "El primer elemento desencolado debe ser 'Hola'")
	require.False(t, cola.EstaVacia(), "La cola vacía debe ser false")
	require.Equal(t, "Mundo!", cola.VerPrimero(), "El elemento primero a ser desencolado ahora es 'Mundo!'")
	require.Equal(t, "Mundo!", cola.Desencolar(), "El elemento desencolado debe ser 'Mundo!'")
	require.True(t, cola.EstaVacia(), "La cola vacía debe ser true")
}

func TestVolumen(t *testing.T) {
	t.Log("Hagamos una prueba de volúmen con 1000 elementos")
	tam := 1000
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 1; i <= tam; i++ {
		cola.Encolar(i)
	}
	require.Equal(t, 1, cola.VerPrimero(), "El primer elemento a desencolar debe ser 1")
	for i := tam; i > 0; i-- {
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia(), "La cola vacía debe ser true")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestCasoIntermitente(t *testing.T) {
	t.Log("Probamos el caso donde se encole y desacole muchas veces")
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		cola.Encolar(i)
		require.Equal(t, cola.VerPrimero(), i, "El primer elemento de la cola no es igual")
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia(), "La cola debe estar vacía luego de terminar todo el proceso")
}
