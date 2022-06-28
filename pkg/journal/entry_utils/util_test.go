package entry_utils

import (
	"testing"
)

func Test_get_prefix_digits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero",
			args: args{0},
			want: 0,
		},
		{
			name: "one",
			args: args{1},
			want: 1,
		},
		{
			name: "ten",
			args: args{10},
			want: 2,
		},
		{
			name: "hundred",
			args: args{100},
			want: 3,
		},
		{
			name: "thousand",
			args: args{1003},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPrefixDigits(tt.args.n); got != tt.want {
				t.Errorf("get_prefix_digits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Fill_prefix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "one",
			args:    args{1},
			want:    "0001",
			wantErr: false,
		},
		{
			name:    "two",
			args:    args{12},
			want:    "0012",
			wantErr: false,
		},
		{
			name:    "three",
			args:    args{123},
			want:    "0123",
			wantErr: false,
		},
		{
			name:    "four",
			args:    args{1234},
			want:    "1234",
			wantErr: false,
		},
		{
			name:    "five",
			args:    args{12345},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FillPrefix(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("fill_prefix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fill_prefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
