package learngo

import (
	"reflect"
)

func percorre(x interface{}, fn func(entrada string)) {

	// versão 1
	// valor := reflect.ValueOf(x) // ValorDe
	// campo := valor.Field(0)     // Campo - equivalente a campo := valor.FieldByName("Nome")
	// fn(campo.String())

	// versão 2
	// valor := reflect.ValueOf(x)

	// for i := 0; i < valor.NumField(); i++ {
	// 	campo := valor.Field(i)
	// 	fn(campo.String())
	// }

	valor := obtemValor(x)

	percorreValor := func(valor reflect.Value) {
		percorre(valor.Interface(), fn)
	}

	switch valor.Kind() {
	case reflect.String:
		fn(valor.String())
	case reflect.Struct:
		for i := 0; i < valor.NumField(); i++ {
			percorreValor(valor.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < valor.Len(); i++ {
			percorreValor(valor.Index(i))
		}
	case reflect.Map:
		for _, chave := range valor.MapKeys() {
			percorreValor(valor.MapIndex(chave))
		}
	}

}

func obtemValor(x interface{}) reflect.Value {
	valor := reflect.ValueOf(x)

	if valor.Kind() == reflect.Ptr {
		valor = valor.Elem()
	}

	return valor
}
