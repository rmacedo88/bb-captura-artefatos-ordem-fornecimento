/**
- Ponto de execução da aplicação
*/
package main

import (
	"bb-captura-artefatos-ordem-fornecimento/CLI"
	"bb-captura-artefatos-ordem-fornecimento/GitLog"
	"bb-captura-artefatos-ordem-fornecimento/Logging"
	"bb-captura-artefatos-ordem-fornecimento/Report"
	"bb-captura-artefatos-ordem-fornecimento/Sys"
	"bufio"
	"os"
)

// TODO Está executando sempre que a ferramenta é acionada. Só deve ser acionado na primeira execução.
func configurarTamanhoHashCommit(comandoGit string) {
	_, _ = Sys.Exec(comandoGit, []string{"config", "--global", "log.abbrevcommit", "yes"}, "")
	_, _ = Sys.Exec(comandoGit, []string{"config", "--global", "core.abbrev", "10"}, "")
}

func carregarRepositorios(arquivoListaRepos string) *[]string {

	arquivo, erroLeitura := os.Open(arquivoListaRepos)

	if erroLeitura != nil {
		Logging.Error.Fatalf("Informe um arquivo válido no parâmetro -arquivo")
		return nil
	}

	leitor := bufio.NewScanner(arquivo)

	leitor.Split(bufio.ScanLines)

	var conteudo []string

	for leitor.Scan() {
		conteudo = append(conteudo, leitor.Text())
	}

	_ = arquivo.Close()

	return &conteudo

}

func logPorRepositorio(parms CLI.Parms, caminhoRepositorioAtual string) *[]byte {

	gitParms := []string{
		"log",
		"--name-status",
		"--author=" + parms.Chave, // chave C
		"--after=" + parms.Data,   // data YYYY-MM-DD
		"--reverse",
		"--pretty=format:#%h",
	}

	caminhoRepositorio := caminhoRepositorioAtual

	resultadoLogGit, erro := Sys.Exec("git", gitParms, caminhoRepositorio)

	if erro != nil {
		Logging.Error.Fatalf(
			"Erro ao capturar o log do repositório %s, verifique os dados e tente novamente.\n",
			caminhoRepositorio,
		)
		return nil
	}

	return &resultadoLogGit
}

func main() {

	Logging.Warning.Println("O uso das informações produzidas por essa ferramenta é de total responsabilidade do utilizador.")

	// Captura os parâmetros informados no terminal
	parms := CLI.Frontend()

	if parms == nil {
		return
	}

	configurarTamanhoHashCommit("git")

	listaRepositorios := carregarRepositorios(parms.Arquivo)

	GitLog.Init()

	for _, repositorioAtual := range *listaRepositorios {

		Logging.Info.Printf("Processando o repositório %s", repositorioAtual)

		resultado := logPorRepositorio(*parms, repositorioAtual)

		if resultado != nil {

			GitLog.SetContext(repositorioAtual)
			_ = GitLog.Parse(string(*resultado))

		} else {
			Logging.Error.Fatalf("Erro ao processar o repositório %s", repositorioAtual)
		}

	}

	// Grava o relatório dos repositórios informados
	Report.Relatorio(GitLog.ListaArtefatos, parms.Chave)
}
