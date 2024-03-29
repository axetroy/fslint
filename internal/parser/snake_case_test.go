package parser

import "testing"

func TestIsSnakeCase(t *testing.T) {
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
		{
			name: "basic",
			args: args{
				str: "do_something2",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "do_something_2",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "do_something_123",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "_do_something_123",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "do_something__2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSnakeCase(tt.args.str); got != tt.want {
				t.Errorf("IsSnakeCase(\"%s\", false) = %v, want %v", tt.args.str, got, tt.want)
			}
		})
	}
}
