/**
- Determina o contrado do artefato
*/
package OrdemFornecimento

import (
	"reflect"
	"testing"
)

func TestBuildNomeArtefato(t *testing.T) {
	type args struct {
		linha string
	}
	tests := []struct {
		name string
		args args
		want Artefato
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := BuildNomeArtefato(tt.args.linha); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. BuildNomeArtefato(%v) = %v, want %v", tt.name, tt.args.linha, got, tt.want)
		}
	}
}

func TestArtefato_String(t *testing.T) {
	type fields struct {
		Nome string
		Tipo string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		artefato := Artefato{
			Nome: tt.fields.Nome,
			Tipo: tt.fields.Tipo,
		}
		if got := artefato.String(); got != tt.want {
			t.Errorf("%q. Artefato.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
