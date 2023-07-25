package service

import (
	"testing"
)

func Test_dummyService_GetVal(t *testing.T) {
	tests := []struct {
		name string
		ds   dummyService
		want string
	}{
		{
			name: "test_service",
			ds:   dummyService{},
			want: "pong",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dummyService{}
			if got := ds.GetVal(); got != tt.want {
				t.Errorf("dummyService.GetVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dummyService_AddVal(t *testing.T) {
	tests := []struct {
		name string
		ds   dummyService
		want string
	}{
		{
			name: "test_addVal",
			ds:   dummyService{},
			want: "something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dummyService{}
			if got := ds.AddVal(); got != tt.want {
				t.Errorf("dummyService.AddVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
