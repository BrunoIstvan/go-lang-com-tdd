package learngo

import (
	"testing"
)

func TestRepetir(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("Repetir quando nao eh passado uma quantidade diferente de zero", func(t *testing.T) {
		repeticoes := Repetir("a", 0)
		esperado := "aaaaa"
		verificaMensagemCorreta(t, repeticoes, esperado)
	})

	t.Run("Repetir quando eh passado uma quantidade especifica de vezes", func(t *testing.T) {
		repeticoes := Repetir("a", 7)
		esperado := "aaaaaaa"
		verificaMensagemCorreta(t, repeticoes, esperado)
	})

}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 8)
	}
}

// func ExampleRepetir() {
// 	repetir := Repetir("a", 4)
// 	fmt.Println(repetir)
// 	// Output: aaaa
// }
