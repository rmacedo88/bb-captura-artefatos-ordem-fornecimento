package Report

import (
	"bb-captura-artefatos-ordem-fornecimento/GitLog"
	"fmt"
	"strings"
)

func Relatorio(listaArtefatos map[string][]GitLog.LinhaLog) {

	for key, dish := range listaArtefatos {
		estado := strings.Split(key, ".")
		fmt.Println()
		fmt.Println(estado[0], estado[1])
		for _, teste := range dish {
			fmt.Println(teste.Projeto + teste.Artefato.String() + teste.Commit)
		}
	}

}
