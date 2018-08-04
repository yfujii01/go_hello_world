package src

import "testing"

func Test_xxx(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name:"test1",want:"あいう えお!!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xxx(); got != tt.want {
				t.Errorf("xxx() = %v, want %v", got, tt.want)
			}
		})
	}
}
