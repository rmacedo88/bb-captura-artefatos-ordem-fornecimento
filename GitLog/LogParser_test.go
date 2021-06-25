package GitLog

import (
	"reflect"
	"testing"
)

func TestSetContext(t *testing.T) {
	type args struct {
		contextoExecucao string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		SetContext(tt.args.contextoExecucao)
	}
}

func TestParse(t *testing.T) {
	type args struct {
		txtLogGit string
	}
	tests := []struct {
		name string
		args args
		want *map[string][]LinhaLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Parse(tt.args.txtLogGit); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Parse() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		Init()
	}
}
