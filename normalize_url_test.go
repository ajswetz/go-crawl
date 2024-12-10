package main

import "testing"

func Test_normalizeURL(t *testing.T) {
	type args struct {
		inputURL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "remove https scheme and trailing slash",
			args: args{
				inputURL: "https://blog.boot.dev/path/",
			},
			want:    "blog.boot.dev/path",
			wantErr: false,
		},
		{
			name: "remove https scheme",
			args: args{
				inputURL: "https://blog.boot.dev/path",
			},
			want:    "blog.boot.dev/path",
			wantErr: false,
		},
		{
			name: "remove http scheme and trailing slash",
			args: args{
				inputURL: "http://blog.boot.dev/path/",
			},
			want:    "blog.boot.dev/path",
			wantErr: false,
		},
		{
			name: "remove http scheme",
			args: args{
				inputURL: "http://blog.boot.dev/path",
			},
			want:    "blog.boot.dev/path",
			wantErr: false,
		},
		{
			name: "no changes required - output should match input",
			args: args{
				inputURL: "blog.boot.dev/path",
			},
			want:    "blog.boot.dev/path",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalizeURL(tt.args.inputURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("normalizeURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("normalizeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
