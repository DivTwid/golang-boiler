package controller

import (
	"reflect"
	"testing"
)

func TestNewDummyController(t *testing.T) {
	tests := []struct {
		name string
		want *DummyController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDummyController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDummyController() = %v, want %v", got, tt.want)
			}
		})
	}
}
