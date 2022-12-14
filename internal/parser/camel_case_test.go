package parser

import "testing"

func TestIsCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "basic",
			args: args{
				str: "DoSomething",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "doSomething",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "do1Something",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "do123Something",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "_do123Something",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "1doSomething",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "Do1something",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "Do1Something",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "Do.Something",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "Do-Something",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "Do_Something",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCamelCase(tt.args.str); got != tt.want {
				t.Errorf("IsCamelCase(\"%s\") = %v, want %v", tt.args.str, got, tt.want)
			}
		})
	}
}
