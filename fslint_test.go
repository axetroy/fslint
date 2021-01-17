package fslint

import (
	"reflect"
	"testing"
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
			name: "basic",
			args: args{
				configFilepath: "fixtures/files.fslint.json",
			},
			wantErr: false,
			want: []LintResult{
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
				{
					FilePath: "fixtures/files/CamelCase11.md",
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
			for _, want := range tt.want {
				catch := false
				for _, g := range got.Values() {
					if g.FilePath == want.FilePath {
						catch = true
						if !reflect.DeepEqual(g, want) {
							t.Errorf("Lint() = %v, want %v", g, want)
						}
					}
				}
				if !reflect.DeepEqual(catch, true) {
					t.Errorf("Lint() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}
