package Sys

import (
	"reflect"
	"testing"
)

func TestExec(t *testing.T) {
	type args struct {
		processo string
		params   []string
		dir      string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Exec(tt.args.processo, tt.args.params, tt.args.dir)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Exec() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Exec() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
