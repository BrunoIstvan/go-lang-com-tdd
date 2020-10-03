package learngo

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	URLLenta := "http://www.facebook.com"
	URLRapida := "http://www.quii.co.uk"

	esperado := URLRapida
	resultado, _ := Corredor(URLLenta, URLRapida)

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func TestExceptionMais10Segundos(t *testing.T) {

	t.Run("retorna um erro se o servidor não responder dentro de 10s", func(t *testing.T) {
		servidorA := criarServidorComAtraso(11 * time.Second)
		servidorB := criarServidorComAtraso(12 * time.Second)

		defer servidorA.Close()
		defer servidorB.Close()

		_, err := Corredor(servidorA.URL, servidorB.URL)

		if err == nil {
			t.Error("esperava um erro, mas não obtive um")
		}
	})

	t.Run("retorna um erro se o servidor não responder dentro de 10s", func(t *testing.T) {
		servidor := criarServidorComAtraso(25 * time.Millisecond)

		defer servidor.Close()

		_, err := Configuravel(servidor.URL, servidor.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("esperava um erro, mas não obtive um")
		}
	})

}

func TestServidoresRapidoELento(t *testing.T) {

	servidorLento := criarServidorComAtraso(20 * time.Millisecond)
	servidorRapido := criarServidorComAtraso(0 * time.Millisecond)

	defer servidorLento.Close()
	defer servidorRapido.Close()

	URLLenta := servidorLento.URL
	URLRapida := servidorRapido.URL

	esperado := URLRapida
	resultado, err := Corredor(URLLenta, URLRapida)

	if err != nil {
		t.Fatalf("não esperava um erro, mas obteve um %v", err)
	}

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}

	t.Run("retorna um erro se o servidor não responder dentro de 10s", func(t *testing.T) {
		servidor := criarServidorComAtraso(25 * time.Millisecond)

		defer servidor.Close()

		_, err := Configuravel(servidor.URL, servidor.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("esperava um erro, mas não obtive um")
		}
	})

}

func criarServidorComAtraso(atraso time.Duration) *httptest.Server {

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		w.WriteHeader(http.StatusOK)
	}))

}
