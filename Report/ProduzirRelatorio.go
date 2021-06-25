/**
 - Grava o relatório em disco
TODO - Consumir os dados de um gerador para aprimorar a performance em múltiplos repositórios
*/
package Report

import (
	"bb-captura-artefatos-ordem-fornecimento/GitLog"
	"bb-captura-artefatos-ordem-fornecimento/Logging"
	"bufio"
	"os"
	"strings"
	"time"
)

func Relatorio(listaArtefatos map[string][]GitLog.LinhaLog, chavec string) {

	mesRef := time.Now()

	nomeArquivo := "relatório-" +
		//strings.TrimSuffix(GitLog.Contexto, string(rune(os.PathSeparator))) +
		chavec +
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

	err := escritor.Flush()

	if err == nil {
		Logging.Info.Println("Relatório produzido com sucesso.")
		Logging.Info.Printf("Confira as informações no arquivo %s.\n", nomeArquivo)
	} else {
		Logging.Error.Fatal("Erro inesperado ao tentar gravar o relatório, por favor tente novamente")
		return
	}

}
