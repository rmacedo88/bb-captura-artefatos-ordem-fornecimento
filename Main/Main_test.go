package main

import "testing"

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
