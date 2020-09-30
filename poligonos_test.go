package learngo

import "testing"

func TestPerimetro(t *testing.T) {

	verificaPerimetro := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()
		resultado := forma.Perimetro()
		if resultado != esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
		}
	}

	t.Run("retangulo", func(t *testing.T) {
		retangulo := Retangulo{Largura: 10.0, Altura: 10.0}
		verificaPerimetro(t, retangulo, 40.0)
	})

	t.Run("circulo", func(t *testing.T) {
		circulo := Circulo{Raio: 10.0}
		verificaPerimetro(t, circulo, 62.83185307179586)
	})

	t.Run("triangulo", func(t *testing.T) {
		triangulo := Triangulo{Base: 10.0, Altura: 5, Lados: [3]float64{10, 10, 10}}
		verificaPerimetro(t, triangulo, 30)
	})

}

func TestArea(t *testing.T) {

	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()
		resultado := forma.Area()
		if resultado != esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
		}
	}

	t.Run("retangulo", func(t *testing.T) {
		retangulo := Retangulo{Largura: 10.0, Altura: 10.0}
		verificaArea(t, retangulo, 100.0)
	})

	t.Run("circulo", func(t *testing.T) {
		circulo := Circulo{Raio: 10.0}
		verificaArea(t, circulo, 314.1592653589793)
	})

	t.Run("triangulo", func(t *testing.T) {
		triangulo := Triangulo{Base: 10.0, Altura: 5, Lados: [3]float64{10, 10, 10}}
		verificaArea(t, triangulo, 25)
	})

}

func TestAreaTableDrivenTests(t *testing.T) {
	testesArea := []struct {
		nome     string
		forma    Forma
		esperado float64
	}{
		{nome: "Retangulo", forma: Retangulo{Largura: 12, Altura: 6}, esperado: 72.0},
		{nome: "Circulo", forma: Circulo{Raio: 10}, esperado: 314.1592653589793},
		{nome: "Triangulo", forma: Triangulo{Base: 10, Altura: 5, Lados: [3]float64{10, 10, 10}}, esperado: 25},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.esperado {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.esperado)
			}
		})
	}

	testesPerimetro := []struct {
		nome     string
		forma    Forma
		esperado float64
	}{
		{nome: "Retangulo", forma: Retangulo{Largura: 10, Altura: 10}, esperado: 40.0},
		{nome: "Circulo", forma: Circulo{Raio: 10}, esperado: 62.83185307179586},
		{nome: "Triangulo", forma: Triangulo{Base: 10, Altura: 5, Lados: [3]float64{10, 10, 10}}, esperado: 30},
	}

	for _, tt := range testesPerimetro {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Perimetro()
			if resultado != tt.esperado {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.esperado)
			}
		})
	}
}
