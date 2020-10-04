package learngo

import (
	"reflect"
	"testing"
)

type Pessoa struct {
	Nome   string
	Perfil Perfil
}

type Perfil struct {
	Idade  int
	Cidade string
}

// func TestPercorre(t *testing.T) {

// 	esperado := "Istvan"
// 	var resultado []string

// 	x := struct {
// 		Nome string
// 	}{esperado}

// 	percorre(x, func(entrada string) {
// 		resultado = append(resultado, entrada)
// 	})

// 	if len(resultado) != 1 {
// 		t.Errorf("número incorreto de chamadas de função: resultado %d, esperado %d", len(resultado), 1)
// 	}

// 	if resultado[0] != esperado {
// 		t.Errorf("resultado '%s', esperado '%s'", resultado[0], esperado)
// 	}

// }

func TestPercorre(t *testing.T) {

	casos := []struct {
		Nome              string
		Entrada           interface{}
		ChamadasEsperadas []string
	}{
		{
			"Struct com um campo string",
			struct {
				Nome string
			}{"Istvan"},
			[]string{"Istvan"},
		},
		{
			"Struct com dois campos tipo string",
			struct {
				Nome   string
				Cidade string
			}{"Istvan", "São Paulo"},
			[]string{"Istvan", "São Paulo"},
		},
		{
			"Struct sem campo tipo string",
			struct {
				Nome  string
				Idade int
			}{"Istvan", 37},
			[]string{"Istvan"},
		},
		{
			"Campos aninhados",
			struct {
				Nome   string
				Perfil struct {
					Idade  int
					Cidade string
				}
			}{"Istvan", struct {
				Idade  int
				Cidade string
			}{37, "São Paulo"}},
			[]string{"Istvan", "São Paulo"},
		},
		{
			"Campos aninhados",
			Pessoa{
				"Istvan",
				Perfil{37, "São Paulo"},
			},
			[]string{"Istvan", "São Paulo"},
		},
		{
			"Ponteiros para coisas",
			&Pessoa{
				"Istvan",
				Perfil{37, "São Paulo"},
			},
			[]string{"Istvan", "São Paulo"},
		},
		{
			"Slices",
			[]Perfil{
				{37, "São Paulo"},
				{34, "Reykjavík"},
			},
			[]string{"São Paulo", "Reykjavík"},
		},
		{
			"Arrays",
			[2]Perfil{
				{37, "São Paulo"},
				{34, "Reykjavík"},
			},
			[]string{"São Paulo", "Reykjavík"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, teste := range casos {

		t.Run(teste.Nome, func(t *testing.T) {

			var resultado []string
			percorre(teste.Entrada, func(entrada string) {
				resultado = append(resultado, entrada)
			})

			if !reflect.DeepEqual(resultado, teste.ChamadasEsperadas) {
				t.Errorf("resultado %v, esperado %v", resultado, teste.ChamadasEsperadas)
			}

		})

	}

	t.Run("com maps", func(t *testing.T) {
		mapA := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var resultado []string
		percorre(mapA, func(entrada string) {
			resultado = append(resultado, entrada)
		})

		verificaSeContem(t, resultado, "Bar")
		verificaSeContem(t, resultado, "Boz")
	})

}

func verificaSeContem(t *testing.T, palheiro []string, agulha string) {
	contem := false
	for _, x := range palheiro {
		if x == agulha {
			contem = true
		}
	}
	if !contem {
		t.Errorf("esperava-se que %+v contivesse '%s', mas não continha", palheiro, agulha)
	}
}
