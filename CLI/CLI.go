package CLI

import (
	"flag"
	"fmt"
)

type CliParms struct {
	Path  string
	Data  string
	Chave string
}

func Frontend() *CliParms {

	path := flag.String("path", "", "Ex.: -path=/home/usuario/kdi/repositorios/sgl-projeto-estatico")
	data := flag.String("data", "", "Ex.: -data=2021-06-01")
	chave := flag.String("chave", "", "Ex.: -chave=C1234567")

	flag.Parse()

	if *path == "" || *data == "" || *chave == "" {
		fmt.Println("Todos os parâmetros são obridatórios, execute novamente informando a opção -h para mais detalhes.")
		return nil
	}

	parms := CliParms{
		Path:  *path,
		Data:  *data,
		Chave: *chave,
	}

	return &parms
}
