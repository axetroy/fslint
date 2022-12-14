package parser

import "testing"

func TestIsKebab(t *testing.T) {
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
			want: true,
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
			want: false,
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
				str: "Do-Staff12",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "Do-Staff-12",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "do-staff12",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "do-staff-12",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "_do-staff-12",
			},
			want: false,
		},
		{
			name: "basic",
			args: args{
				str: "do-staff--12",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKebabCase(tt.args.str); got != tt.want {
				t.Errorf("IsKebabCase(\"%s\") = %v, want %v", tt.args.str, got, tt.want)
			}
		})
	}
}
