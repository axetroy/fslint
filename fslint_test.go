package fslint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLint(t *testing.T) {
	type args struct {
		configFilepath string
	}
	tests := []struct {
		name    string
		args    args
		want    []LintResult
		wantErr bool
	}{
		{
			name: "files",
			args: args{
				configFilepath: "fixtures/files.fslintrc.json",
			},
			wantErr: false,
			want: []LintResult{
				{
					FilePath: "fixtures/files/CamelCase11.md",
					Expect:   ModeLittleKebab,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/Kebab-Kebab21.md",
					Expect:   ModeLittleKebab,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/Snake_Case31.md",
					Expect:   ModeLittleKebab,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/camelCase12.md",
					Expect:   ModeLittleKebab,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/snake_case32.md",
					Expect:   ModeLittleKebab,
					Level:    LevelError,
				},
			},
		},
		{
			name: "folder",
			args: args{
				configFilepath: "fixtures/folder.fslintrc.json",
			},
			wantErr: false,
			want: []LintResult{
				{
					FilePath: "fixtures/TestFolder",
					Expect:   ModeLittleKebab,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/TestFolder/Nest_Folder",
					Expect:   ModeLittleKebab,
					Level:    LevelError,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lint(tt.args.configFilepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got.Values(), tt.name)
		})
	}
}
