package entry_utils

import (
	"reflect"
	"testing"
)

func Test_trim_newlines(t *testing.T) {
	type args struct {
		editor_contents []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "no newlines",
			args: args{[]byte("hello")},
			want: []byte("hello"),
		},
		{
			name: "one newline",
			args: args{[]byte("hello\n")},
			want: []byte("hello"),
		},
		{
			name: "multiple newlines",
			args: args{[]byte("hello\n\n\n\n\n")},
			want: []byte("hello"),
		},
		{
			name: "empty string",
			args: args{[]byte("")},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trim_newlines(tt.args.editor_contents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trim_newlines() = %v, want %v", got, tt.want)
			}
		})
	}
}
