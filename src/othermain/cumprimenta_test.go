// package main
package othermain

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {

	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Bruno")

	resultado := buffer.String()
	esperado := "Olá, Bruno"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
