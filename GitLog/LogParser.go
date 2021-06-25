/**
- Efetua o parsing do log obtido do git
*/
package GitLog

import (
	"bb-captura-artefatos-ordem-fornecimento/OrdemFornecimento"
	"os"
	"strings"
)

// Contexto (nome do repositório)
var Contexto string

// Guarda os artefatos que serão usados no relatório final
var ListaArtefatos map[string][]LinhaLog // := make(map[string][]LinhaLog)

// Inicializa o mapa que deverá guardar os artefatos parseados para gerar somente um relatório
func Init() {
	ListaArtefatos = make(map[string][]LinhaLog)
}

// Determina o nome do contexto de execução - diretório onde o repositório da vez está contido
func SetContext(contextoExecucao string) {
	// Recupera a posição do último separador de arquivo
	ultimoSeparador := strings.LastIndexByte(contextoExecucao, os.PathSeparator)
	// Extrai o nome do projeto
	Contexto = contextoExecucao[ultimoSeparador+1:] + string(os.PathSeparator)
}

/**
- Efetua o parsing do log obtido do repositório
- O mapa retornado está depreciado, seu valor é acumulado na variável @ListaArtefatos
*/
func Parse(txtLogGit string) *map[string][]LinhaLog {

	// Extrai as linhas do log obtido do repositório
	linhas := strings.Split(txtLogGit, "\n")

	if len(linhas) == 0 {
		return nil
	}

	// Byte que identifica os artefatos removidos
	REMOCAO := OrdemFornecimento.EstadoArtefatoFromString("D").Index()

	// Guarda os artefatos removidos para evitar a inclusão destes no relatório final
	artefatosRemovidos := make(map[string][]string)

	// Armazena o hash do último commit
	ultimoCommit := ""

	for _, linhaLog := range linhas {

		// Constrói o objeto com as informações da linha atual
		linha := buildLinhaLog(linhaLog)

		if linha != nil {
			// Atualiza o último commit
			if len(linha.Commit) > 1 {
				ultimoCommit = linha.Commit
			} else {
				// Copia o último commit para a linha atual
				linha.Commit = ultimoCommit

				if linha.EstadoAtf.Index() == REMOCAO {
					artefatosRemovidos[linha.Artefato.String()] =
						append(artefatosRemovidos[linha.Artefato.String()], linha.Artefato.String())
				}

				_, itemExcluido := artefatosRemovidos[linha.Artefato.String()]

				if !itemExcluido {
					ListaArtefatos[string((linha.EstadoAtf.String()))+"."+linha.Artefato.Tipo] =
						append(ListaArtefatos[string((linha.EstadoAtf.String()))+"."+linha.Artefato.Tipo], *linha)
				}
			}
		}

	}

	return &ListaArtefatos

}
