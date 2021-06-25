package main

import (
	"bb-captura-artefatos-ordem-fornecimento/CLI"
	"reflect"
	"testing"
)

func Test_configurarTamanhoHashCommit(t *testing.T) {
	type args struct {
		comandoGit string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		configurarTamanhoHashCommit(tt.args.comandoGit)
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		main()
	}
}

func Test_carregarRepositorios(t *testing.T) {
	type args struct {
		arquivoListaRepos string
	}
	tests := []struct {
		name string
		args args
		want *[]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := carregarRepositorios(tt.args.arquivoListaRepos); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. carregarRepositorios(%v) = %v, want %v", tt.name, tt.args.arquivoListaRepos, got, tt.want)
		}
	}
}

func Test_logPorRepositorio(t *testing.T) {
	type args struct {
		parms                   CLI.Parms
		caminhoRepositorioAtual string
	}
	tests := []struct {
		name string
		args args
		want *[]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := logPorRepositorio(tt.args.parms, tt.args.caminhoRepositorioAtual); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. logPorRepositorio(%v, %v) = %v, want %v", tt.name, tt.args.parms, tt.args.caminhoRepositorioAtual, got, tt.want)
		}
	}
}
