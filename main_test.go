package main

import "testing"

func Test_kakezan(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name:"t1",want:42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kakezan(); got != tt.want {
				t.Errorf("kakezan() = %v, want %v", got, tt.want)
			}
		})
	}
}
