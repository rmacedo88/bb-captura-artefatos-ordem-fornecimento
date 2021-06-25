/**
- Grava o relatório em disco
*/
package Report

import (
	"bb-captura-artefatos-ordem-fornecimento/GitLog"
	"testing"
)

func TestRelatorio(t *testing.T) {
	type args struct {
		listaArtefatos map[string][]GitLog.LinhaLog
		chavec         string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Relatorio(tt.args.listaArtefatos, tt.args.chavec)
	}
}
