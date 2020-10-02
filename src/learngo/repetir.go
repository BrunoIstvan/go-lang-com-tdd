package learngo

const quantidadeRepeticoes = 5

/*
	Repete o caracter quantas vezes for informado no parametro quantidade
*/
func Repetir(caracter string, quantidade int) string {

	if quantidade == 0 {
		quantidade = quantidadeRepeticoes
	}

	var repeticoes string

	for i := 0; i < quantidade; i++ {
		repeticoes += caracter
	}

	return repeticoes

}
