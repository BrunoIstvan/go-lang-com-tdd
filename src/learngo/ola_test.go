package learngo

import "testing"

// func TestOla(t *testing.T) {
// 	resultado := Ola()
// 	esperado := "Olá, mundo"

// 	if resultado != esperado {
// 		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
// 	}
// }

// func TestOla(t *testing.T) {
// 	resultado := Ola("Bruno")
// 	esperado := "Olá, Bruno"

// 	if resultado != esperado {
// 		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
// 	}
// }

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris", "")
		esperado := "Olá, Chris"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", espanhol)
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := Ola("Zinedine", frances)
		esperado := "Bonjour, Zinedine"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em inglês", func(t *testing.T) {
		resultado := Ola("Vardy", ingles)
		esperado := "Hi, Vardy"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	// t.Run("em alemão", func(t *testing.T) {
	// 	resultado := Ola("Muller", alemao)
	// 	esperado := "Hi, Muller"
	// 	verificaMensagemCorreta(t, resultado, esperado)
	// })

}
