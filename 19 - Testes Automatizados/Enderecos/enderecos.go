package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
// Exemplo de uso: TipoDeEndereco("Rua das Flores")
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetrasMinusculas := cases.Lower(language.BrazilianPortuguese).String(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetrasMinusculas, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
			break
		}
	}

	if enderecoTemUmTipoValido {
		return cases.Title(language.BrazilianPortuguese).String(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
