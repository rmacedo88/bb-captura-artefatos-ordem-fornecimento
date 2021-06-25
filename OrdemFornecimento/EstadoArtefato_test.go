package OrdemFornecimento

import "testing"

func TestEstadoArtefato_String(t *testing.T) {
	tests := []struct {
		name       string
		estadoEnum EstadoArtefato
		want       string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.estadoEnum.String(); got != tt.want {
			t.Errorf("%q. EstadoArtefato.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestEstadoArtefato_Index(t *testing.T) {
	tests := []struct {
		name       string
		estadoEnum EstadoArtefato
		want       byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.estadoEnum.Index(); got != tt.want {
			t.Errorf("%q. EstadoArtefato.Index() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestEstadoArtefatoFromString(t *testing.T) {
	type args struct {
		literal string
	}
	tests := []struct {
		name string
		args args
		want EstadoArtefato
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := EstadoArtefatoFromString(tt.args.literal); got != tt.want {
			t.Errorf("%q. EstadoArtefatoFromString() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
