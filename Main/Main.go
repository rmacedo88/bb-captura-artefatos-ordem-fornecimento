/**
- Ponto de execução da aplicação
*/
package main

import (
	"bb-captura-artefatos-ordem-fornecimento/CLI"
	"bb-captura-artefatos-ordem-fornecimento/GitLog"
	"bb-captura-artefatos-ordem-fornecimento/Report"
	"bb-captura-artefatos-ordem-fornecimento/Sys"
)

func configurarTamanhoHashCommit(comandoGit string) {
	_, _ = Sys.Exec(comandoGit, []string{"config", "--global", "log.abbrevcommit", "yes"}, "")
	_, _ = Sys.Exec(comandoGit, []string{"config", "--global", "core.abbrev", "10"}, "")
}

func main() {

	// Captura os parâmetros informados no terminal
	parms := CLI.Frontend()

	gitCommand := "git"

	configurarTamanhoHashCommit(gitCommand)

	gitParms := []string{
		"log",
		"--name-status",
		"--author=" + parms.Chave, // chave C
		"--after=" + parms.Data,   // data YYYY-MM-DD
		"--reverse",
		"--pretty=format:#%h",
	}

	contextoExecucao := parms.Path

	resultado, erro := Sys.Exec(gitCommand, gitParms, contextoExecucao)

	if erro != nil {
		println(erro, "Erro ao tentar capturar o log do repositório informado")
	}

	GitLog.SetContext(contextoExecucao)
	entrega := GitLog.Parse(string(resultado))

	Report.Relatorio(*entrega)
}

/**
//go run Main.go -chave=C1218068 -data=2019-06-01 -path=/media/dev/af397934-0a17-4e0f-9d50-d5146aac487c/home/robson/repos-teste/stt-fluxo-entrega-estatico

*/
