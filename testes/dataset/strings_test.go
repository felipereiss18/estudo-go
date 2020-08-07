package strings

import (
	"strings"
	"testing"
)

const msgError = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)"

func TestIndex(t *testing.T)  {
	testes := []struct{
		texto string
		parte string
		esperado int
	}{
		{"Go (Golang)", "lang", 6},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Felipe", "e", 1},
	}

	for _, teste := range testes {
		t.Logf("Dado testado: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgError, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
