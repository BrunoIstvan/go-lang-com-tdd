package learngo

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%d', esperado '%d'", resultado, esperado)
		}
	}

	t.Run("soma simples", func(t *testing.T) {
		soma := Soma(2, 2)
		esperado := 4
		verificaMensagemCorreta(t, soma, esperado)
	})
}

func TestSomaSlice(t *testing.T) {

	t.Run("soma de numeros de um slice", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}
		resultado := SomaSlice(numeros)
		esperado := 15
		if resultado != esperado {
			t.Errorf("resultado '%d', esperado '%d', dado '%v'", resultado, esperado, numeros)
		}
	})

}

func TestSomaArray(t *testing.T) {

	t.Run("soma de numeros de um array", func(t *testing.T) {
		numeros := [5]int{1, 2, 3, 5, 6}
		resultado := SomaArray(numeros)
		esperado := 17
		if resultado != esperado {
			t.Errorf("resultado '%d', esperado '%d', dado '%v'", resultado, esperado, numeros)
		}
	})

}

func TestSomaTudo(t *testing.T) {

	t.Run("Teste simples de slice", func(t *testing.T) {
		recebido := SomaTudo([]int{1, 2}, []int{0, 9})
		esperado := []int{3, 9}

		if !reflect.DeepEqual(recebido, esperado) {
			t.Errorf("resultado %v esperado %v", recebido, esperado)
		}
	})

	t.Run("Testando com 2 slices", func(t *testing.T) {
		recebido := SomaTudo([]int{1, 2, 3, 4, 5}, []int{0, 9, 10})
		esperado := []int{15, 19}

		if !reflect.DeepEqual(recebido, esperado) {
			t.Errorf("resultado %v esperado %v", recebido, esperado)
		}
	})

	t.Run("Testando com 4 slices", func(t *testing.T) {
		recebido := SomaTudo([]int{1, 2, 3, 4, 5}, []int{0, 9, 10}, []int{1, 1, 1}, []int{10, 19, 10})
		esperado := []int{15, 19, 3, 39}

		if !reflect.DeepEqual(recebido, esperado) {
			t.Errorf("resultado %v esperado %v", recebido, esperado)
		}
	})

}

func TestSomaTodoOResto(t *testing.T) {

	verificaSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v, esperado %v", resultado, esperado)
		}
	}

	t.Run("faz a soma do resto", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}
		verificaSomas(t, resultado, esperado)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}
		verificaSomas(t, resultado, esperado)
	})

}

// func ExampleSomaSimples() {
// 	soma := Soma(1, 5)
// 	fmt.Println(soma)
// 	// Output: 6
// }

// func ExampleSomaSlice() {
// 	soma := SomaSlice([]int{1, 3, 5})
// 	fmt.Println(soma)
// 	// Output: 9
// }

// func ExampleSomaFixo() {
// 	soma := SomaArray([5]int{2, 2, 2, 3, 5})
// 	fmt.Println(soma)
// 	// Output: 14
// }
