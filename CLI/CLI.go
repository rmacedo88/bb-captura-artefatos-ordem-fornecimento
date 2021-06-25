/**
- Interface de linha de comando permitindo ajustar a chave, data e diretório de trabalho da OF
*/

package CLI

import (
	"bb-captura-artefatos-ordem-fornecimento/Logging"
	"flag"
)

type Parms struct {
	Arquivo string
	Data    string
	Chave   string
}

func Frontend() *Parms {

	arquivo := flag.String("arquivo", "", "Ex.: -arquivo=arquivo_com_a_lista_de_repositorios.txt")
	data := flag.String("data", "", "Ex.: -data=2021-06-01")
	chave := flag.String("chave", "", "Ex.: -chave=C1234567")

	flag.Parse()

	// Valida os parâmetros - todos são obrigatórios
	if *arquivo == "" || *data == "" || *chave == "" {
		Logging.Error.Fatal("Todos os parâmetros são obrigatórios, execute novamente informando a opção -h para mais detalhes.")
		Logging.Error.Println("Dúvidas? https://github.com/rmacedo88/bb-captura-artefatos-ordem-fornecimento#readme")
		return nil
	}

	parms := Parms{
		Arquivo: *arquivo,
		Data:    *data,
		Chave:   *chave,
	}

	return &parms
}
