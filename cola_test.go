package cola_test

import (
	TDACola "cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia(), "La cola vacía deber ser true")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestUnElemento(t *testing.T) {
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
