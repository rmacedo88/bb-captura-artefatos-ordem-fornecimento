/**
- Converte um índice do mês do ano para um texto legível
*/
package Report

import "testing"

func TestMes_String(t *testing.T) {
	tests := []struct {
		name       string
		estadoEnum Mes
		want       string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.estadoEnum.String(); got != tt.want {
			t.Errorf("%q. Mes.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
