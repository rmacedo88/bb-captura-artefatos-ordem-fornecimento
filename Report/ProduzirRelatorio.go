/**
 - Grava o relatório em disco
TODO - Consumir os dados de um gerador para aprimorar a performance em múltiplos repositórios
*/
package Report

import (
	"bb-captura-artefatos-ordem-fornecimento/GitLog"
	"bufio"
	"os"
	"strings"
	"time"
)

func Relatorio(listaArtefatos map[string][]GitLog.LinhaLog) {

	mesRef := time.Now()

	nomeArquivo := "relatório-" +
		strings.TrimSuffix(GitLog.Contexto, string(rune(os.PathSeparator))) +
		"-" +
		Mes(mesRef.Month()).String() +
		".txt"

	diretorioAtual, _ := os.Getwd()

	arquivoRelatorio, _ := os.Create(diretorioAtual + string(os.PathSeparator) + nomeArquivo)

	escritor := bufio.NewWriter(arquivoRelatorio)

	// Itera sobre as informações para gravar o arquivo

	for chave, lista := range listaArtefatos {

		titulo := strings.Split(chave, ".")

		_, _ = escritor.WriteString("\n")
		_, _ = escritor.WriteString(titulo[0] + " " + titulo[1] + "\n")

		for _, linhaRelatorio := range lista {
			_, _ = escritor.WriteString(linhaRelatorio.Projeto +
				linhaRelatorio.Artefato.String() +
				linhaRelatorio.Commit + "\n")
		}
	}

	_ = escritor.Flush()

}
