package OrdemFornecimento

import (
	"bytes"
	"strings"
)

type Artefato struct {
	Nome string // Nome do artefato sem a extensão, tal como ocorre no log do git
	Tipo string // Extensão do artefato
}

/**
- Dada uma linha de log do git, devolve o artefato
*/
func BuildNomeArtefato(linha string) Artefato {

	artefato := Artefato{}

	// Recupera a posição do último ponto
	ultimoPonto := strings.LastIndexByte(linha, '.')

	// Se houver um único ponto, devolve somente o tipo
	if ultimoPonto == -1 {
		artefato.Nome = ""
		artefato.Tipo = linha

		return artefato
	}

	// Efetua o build do artefato
	artefato.Nome = linha[:ultimoPonto]
	artefato.Tipo = linha[ultimoPonto+1:]

	return artefato
}

func (artefato Artefato) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(artefato.Nome)
	buffer.WriteString(".")
	buffer.WriteString(artefato.Tipo)

	return buffer.String()
}
