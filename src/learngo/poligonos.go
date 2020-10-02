package learngo

import "math"

// Definindo Forma como interface e declarando os metodos que devem ser implementados
type Forma interface {
	Area() float64
	Perimetro() float64
}

/*
	Como declaramos que Retangulo tem os mesmos metodos da interface,
 	a vinculacao eh feita automaticamente,
 	logo Retangulo passa a ser uma Forma
*/
type Retangulo struct {
	Largura float64
	Altura  float64
}

// Calcula Perimetro de um Retangulo
func (retangulo Retangulo) Perimetro() float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

// Calcula Area de um Retangulo
func (retangulo Retangulo) Area() float64 {
	return retangulo.Largura * retangulo.Altura
}

/*
	Como declaramos que Circulo tem os mesmos metodos da interface,
 	a vinculacao eh feita automaticamente,
 	logo Retangulo passa a ser uma Forma
*/
type Circulo struct {
	Raio float64
}

// Calcula Perimetro de um Circulo
func (circulo Circulo) Perimetro() float64 {
	return 2 * math.Pi * circulo.Raio
}

// Calcula Area de um Circulo
func (circulo Circulo) Area() float64 {
	return math.Pi * circulo.Raio * circulo.Raio
}

/*
	Como declaramos que Triangulo tem os mesmos metodos da interface,
 	a vinculacao eh feita automaticamente,
 	logo Retangulo passa a ser uma Forma
*/
type Triangulo struct {
	Base   float64
	Altura float64
	Lados  [3]float64
}

// Calcula Perimetro de um Triangulo
func (triangulo Triangulo) Perimetro() float64 {
	var perimetro float64
	for _, lado := range triangulo.Lados {
		perimetro += lado
	}
	return perimetro
}

// Calcula Area de um Triangulo
func (triangulo Triangulo) Area() float64 {
	return (triangulo.Base * triangulo.Altura) / 2
}
