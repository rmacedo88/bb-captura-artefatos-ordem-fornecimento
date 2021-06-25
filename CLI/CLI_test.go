package CLI

import (
	"reflect"
	"testing"
)

func TestFrontend(t *testing.T) {
	var tests []struct {
		name string
		want *Parms
	}
	for _, tt := range tests {
		if got := Frontend(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Frontend() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
