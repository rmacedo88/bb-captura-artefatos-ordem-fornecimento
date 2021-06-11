package Sys

import (
	"bytes"
	"os/exec"
)

/**
- Executa um processo dados nome, parâmetros e contexto de execução
*/
func Exec(processo string, params []string, dir string) ([]byte, error) {

	// Inicia o processo
	cmd := exec.Command(processo, params...)

	// Cria os objetos que receberão a saída padrão e de erros
	var stdout, stderr bytes.Buffer

	// Atribui os ponteiros à respectiva saída normal e de erro do comando
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Define o contexto de execução do processo
	if dir != "" {
		cmd.Dir = dir
	}

	// Aciona o processo
	err := cmd.Run()

	// Devolve os erros caso existam
	if err != nil {
		return nil, err
	}

	// Devolve a saída padrão do processo executado
	return stdout.Bytes(), nil
}
