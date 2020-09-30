package learngo

/*
	Método que retorna a soma de dois numeros passados por parametro
*/
func Soma(x, y int) int {
	return x + y
}

/*
	Método que retorna a soma dos numeros dentro do array passado por parametro
*/
func SomaArray(numeros [5]int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

/*
	Método que retorna a soma dos numeros dentro do slice passado por parametro
*/
func SomaSlice(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func SomaTudo(numerosParaSomar ...[]int) (somas []int) {

	for _, numeros := range numerosParaSomar {
		somas = append(somas, SomaSlice(numeros))
	}

	return
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, SomaSlice(final))
		}
	}

	return somas
}
