package db

import "testing"

func TestInitDbModel(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "case",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitDbModel(); (err != nil) != tt.wantErr {
				t.Errorf("InitDbModel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
