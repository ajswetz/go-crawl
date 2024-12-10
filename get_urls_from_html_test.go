package main

import (
	"reflect"
	"testing"
)

func Test_getURLsFromHTML(t *testing.T) {

	htmlSample := `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`

	type args struct {
		htmlBody   string
		rawBaseURL string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "absolute and relative URLs in one sample",
			args: args{
				htmlBody:   htmlSample,
				rawBaseURL: "https://boot.dev",
			},
			want: []string{"https://boot.dev/path/one", "https://other.com/path/one"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getURLsFromHTML(tt.args.htmlBody, tt.args.rawBaseURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("getURLsFromHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getURLsFromHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
