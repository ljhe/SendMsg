package log

import "testing"

func TestErr(t *testing.T) {
	type args struct {
		str   string
		param []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case",
			args: args{
				str:   "test %d %v",
				param: []interface{}{777, "777"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Err(tt.args.str, tt.args.param...)
		})
	}
}
