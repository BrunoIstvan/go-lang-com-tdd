package learngo

import (
	"fmt"
)

const portugues = "portugues"
const espanhol = "espanhol"
const frances = "francês"
const ingles = "inglês"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaIngles = "Hi, "

/*
	Metodo que retorna um ola mundo ou ola personalizado para alguem em idioma definido
*/
func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {

	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	case ingles:
		prefixo = prefixoOlaIngles
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
