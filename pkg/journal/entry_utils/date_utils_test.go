package entry_utils

import (
	"fmt"
	"testing"
	"time"
)

func Test_get_days_back(t *testing.T) {
	type args struct {
		day string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Monday",
			args: args{fmt.Sprint(time.Monday)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := get_days_back(tt.args.day); got != tt.want {
				t.Errorf("get_days_back() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_get_month_from_string(t *testing.T) {
	type args struct {
		month string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Month
		wantErr bool
	}{
		{
			name: "February",
			args: args{"February"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := get_month_from_string(tt.args.month)
			if (err != nil) != tt.wantErr {
				t.Errorf("get_month_from_string() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("get_month_from_string() = %v, want %v", got, tt.want)
			}
		})
	}
}
