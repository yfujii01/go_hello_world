package src

import "testing"

func Test_bar(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "t1", args: args{s1: "asd", s2: "fgh"}, want: "asdfgh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bar(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("bar() = %v, want %v", got, tt.want)
			}
		})
	}
}
