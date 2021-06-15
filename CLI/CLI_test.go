package CLI

import (
	"reflect"
	"testing"
)

func TestFrontend(t *testing.T) {
	tests := []struct {
		name string
		want *CliParms
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Frontend(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Frontend() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
