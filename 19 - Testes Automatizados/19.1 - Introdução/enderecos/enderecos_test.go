package enderecos

import "testing"

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoEndereco(t *testing.T) {

	eventosTeste := []cenarioTeste{
		{"Rua ABC", "rua"},
		{"Avenida Brasil", "avenida"},
		{"Estrada do Primavera", "estrada"},
		{"Rodovia da BR 364", "rodovia"},
		{"Escritório da Maringá", "Tipo Inválido"},
		{"rua LALALALA", "rua"},
		{"AVENIDA XOXOXO", "avenida"},
	}

	for _, cenario := range eventosTeste {
		tipoDeEnderecoRecebido := TipoEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido, cenario.retornoEsperado,
			)
		}
	}
}
