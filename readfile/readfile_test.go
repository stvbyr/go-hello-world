package readfile

import "testing"

func TestReadMarkdownFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"File Present", args{filename: "./testfile.md"}, "# Testfile", false},
		{"File not Present", args{filename: "./no-there.md"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadMarkdownFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadMarkdownFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadMarkdownFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
