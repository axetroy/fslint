package parser

import "testing"

func TestIsBigSnakeCase(t *testing.T) {
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
				str: "Do-something",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "do-something",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "Do_Something",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "do_something",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "Dostaff",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "dostaff",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSnakeCase(tt.args.str, true); got != tt.want {
				t.Errorf("IsBigSnakeCase(\"%s\", true) = %v, want %v", tt.args.str, got, tt.want)
			}
		})
	}
}

func TestIsLittleSnakeCase(t *testing.T) {
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
				str: "Do-something",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "do-something",
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
		{
			name: "basic",
			args: args{
				str: "do_something",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "Dostaff",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "dostaff",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSnakeCase(tt.args.str, false); got != tt.want {
				t.Errorf("IsLittleSnakeCase(\"%s\", false) = %v, want %v", tt.args.str, got, tt.want)
			}
		})
	}
}
