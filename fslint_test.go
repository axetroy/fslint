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
					Expect:   ModeKebabCase,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/Kebab-Kebab21.md",
					Expect:   ModeKebabCase,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/Snake_Case31.md",
					Expect:   ModeKebabCase,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/camelCase12.md",
					Expect:   ModeKebabCase,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/snake_case32.md",
					Expect:   ModeKebabCase,
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
					Expect:   ModeKebabCase,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/TestFolder/Nest_Folder",
					Expect:   ModeKebabCase,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/TestFolder/ignore_folder",
					Expect:   ModeKebabCase,
					Level:    LevelError,
				},
			},
		},
		{
			name: "ignore files",
			args: args{
				configFilepath: "fixtures/ignore-files.fslintrc.json",
			},
			wantErr: false,
			want: []LintResult{
				{
					FilePath: "fixtures/files/CamelCase11.md",
					Expect:   ModeKebabCase,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/Kebab-Kebab21.md",
					Expect:   ModeKebabCase,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/camelCase12.md",
					Expect:   ModeKebabCase,
					Level:    LevelError,
				},
			},
		},
		{
			name: "ignore folder",
			args: args{
				configFilepath: "fixtures/ignore-folder.fslintrc.json",
			},
			wantErr: false,
			want: []LintResult{
				{
					FilePath: "fixtures/TestFolder",
					Expect:   ModeKebabCase,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/TestFolder/Nest_Folder",
					Expect:   ModeKebabCase,
					Level:    LevelError,
				},
			},
		},
		{
			name: "files",
			args: args{
				configFilepath: "fixtures/regexp.fslintrc.json",
			},
			wantErr: false,
			want: []LintResult{
				{
					FilePath: "fixtures/files/CamelCase11.md",
					Expect:   ModeRegExp,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/Kebab-Kebab21.md",
					Expect:   ModeRegExp,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/Snake_Case31.md",
					Expect:   ModeRegExp,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/camelCase12.md",
					Expect:   ModeRegExp,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/kebab-kebab22.md",
					Expect:   ModeRegExp,
					Level:    LevelError,
				},
				{
					FilePath: "fixtures/files/snake_case32.md",
					Expect:   ModeRegExp,
					Level:    LevelError,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lint(tt.args.configFilepath, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got.Values(), tt.name)
		})
	}
}
