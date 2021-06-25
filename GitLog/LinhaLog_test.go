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
			t.Errorf("%q. buildLinhaLog() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
