package service

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestHandleRawData(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleRawData()
		})
	}
}
