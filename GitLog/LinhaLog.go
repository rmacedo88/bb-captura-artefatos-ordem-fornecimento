/**
- Efetua o parsing de uma linha obtida do log do git
*/
package GitLog

import (
	"bb-captura-artefatos-ordem-fornecimento/OrdemFornecimento"
	"strings"
)

type LinhaLog struct {
	EstadoAtf OrdemFornecimento.EstadoArtefato // criação, alteração ou remoção
	Projeto   string                           // Artefato do projeto onde o artefato está inserido
	Artefato  OrdemFornecimento.Artefato       // Nome e tipo do artefato
	Commit    string                           // hash do commit
}

func buildLinhaLog(linha string) *LinhaLog {

	if len(linha) == 0 {
		return nil
	}

	linhaLog := LinhaLog{}

	// Extrai o commit hash do commit e retorna
	if strings.HasPrefix(linha, "#") {
		linhaLog.Commit = linha // strings.TrimFunc(linha, func(r rune) bool { return unicode.IsPunct(r) })
		return &linhaLog
	}

	// Quebra a string pela tabulação, resultando no estado do artefato e o nome
	dadosLinha := strings.Split(linha, "\t")

	// Extrai o estado, artefato e retorna
	if len(dadosLinha) == 2 {
		linhaLog.EstadoAtf = OrdemFornecimento.EstadoArtefatoFromString(dadosLinha[0])
		linhaLog.Artefato = OrdemFornecimento.BuildNomeArtefato(dadosLinha[1])
		linhaLog.Projeto = Contexto
		return &linhaLog
	}

	return nil
}
