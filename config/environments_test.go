package config

import "testing"

func TestLoadENV(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{

			name: "test1",
			args: args{
				path: "../env/.env_test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoadENV(tt.args.path)
		})
	}
}
