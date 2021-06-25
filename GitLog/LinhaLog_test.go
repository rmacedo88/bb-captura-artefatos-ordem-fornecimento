/**
- Efetua o parsing de uma linha obtida do log do git
*/
package GitLog

import (
	"reflect"
	"testing"
)

func Test_buildLinhaLog(t *testing.T) {
	type args struct {
		linha string
	}
	tests := []struct {
		name string
		args args
		want *LinhaLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := buildLinhaLog(tt.args.linha); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. buildLinhaLog(%v) = %v, want %v", tt.name, tt.args.linha, got, tt.want)
		}
	}
}
