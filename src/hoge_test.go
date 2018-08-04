package src

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "test1", want: "Hello World!!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinMessage(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "ジョインテスト1", args: args{s1: "aa", s2: "bb"}, want: "aa bb!!"},
		{name: "ジョインテスト2", args: args{s1: "bb", s2: "cc"}, want: "bb cc!!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinMessage(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("JoinMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
