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
				configFilepath: ".fslint.json",
			},
			wantErr: false,
			want: []LintResult{
				{
					FilePath: "__test__/Kebab-Kebab21.md",
					FileName: "Kebab-Kebab21.md",
					Expect:   ModeLittleKebab,
				},
				{
					FilePath: "__test__/Snake_Case31.md",
					FileName: "Snake_Case31.md",
					Expect:   ModeLittleKebab,
				},
				{
					FilePath: "__test__/camelCase12.md",
					FileName: "camelCase12.md",
					Expect:   ModeLittleKebab,
				},
				{
					FilePath: "__test__/snake_case32.md",
					FileName: "snake_case32.md",
					Expect:   ModeLittleKebab,
				},
				{
					FilePath: "__test__/CamelCase11.md",
					FileName: "CamelCase11.md",
					Expect:   ModeLittleKebab,
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
				for _, g := range got {
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
