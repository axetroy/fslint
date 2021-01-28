package parser

import "testing"

func TestIsBigDotDot(t *testing.T) {
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
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "do.something",
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
				str: "Do.1.Staff",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "Do.123.Staff",
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
			if got := IsDotDot(tt.args.str, true); got != tt.want {
				t.Errorf("IsBigDotDot(\"%s\", true) = %v, want %v", tt.args.str, got, tt.want)
			}
		})
	}
}

func TestIsLittleDotDot(t *testing.T) {
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
				str: "do.something",
			},
			want: true,
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
				str: "do.staff",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "do.1.staff",
			},
			want: true,
		},
		{
			name: "basic",
			args: args{
				str: "do.123.staff",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDotDot(tt.args.str, false); got != tt.want {
				t.Errorf("IsLittleDotDot(\"%s\", false) = %v, want %v", tt.args.str, got, tt.want)
			}
		})
	}
}
