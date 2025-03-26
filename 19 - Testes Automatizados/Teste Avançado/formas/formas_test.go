package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{
			Altura: 10, Largura: 12,
		}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{
			raio: 10,
		}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()
		if areaEsperada != areaRecebida {
			t.Errorf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
