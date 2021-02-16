package main

import (
"testing"
)

func TestReverseString1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing empty",
			args: args{
				s:"",
			},
			want: "",
		},
		{
			name: "testing hello",
			args: args{
				s:"hello",
			},
			want: "olleh",
		},
		{
			name: "testing question",
			args: args{
				s:"why am I?",
			},
			want: "?I ma yhw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}