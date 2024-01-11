package parenthesisremoval

import "testing"

func TestParenthesisRemoval(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				expr: "(a-(b-c))",
			},
			want: "a-(b-c)",
		},
		{
			name: "2",
			args: args{
				expr: "(a*(b+c))",
			},
			want: "a*(b+c)",
		},
		{
			name: "3",
			args: args{
				expr: "(a+(b+c))",
			},
			want: "a+b+c",
		},
		{
			name: "4",
			args: args{
				expr: "a+(b+c)*d",
			},
			want: "a+(b+c)*d",
		},
		{
			name: "5",
			args: args{
				expr: "(a-b)+c*d",
			},
			want: "a-b+c*d",
		},
		{
			name: "6",
			args: args{
				expr: "(((a)))",
			},
			want: "a",
		},
		{
			name: "7",
			args: args{
				expr: "(((a+b)))",
			},
			want: "a+b",
		},
		{
			name: "8",
			args: args{
				expr: "(c)*(((a+b)))",
			},
			want: "c*(a+b)",
		},
		{
			name: "9",
			args: args{
				expr: "-(2)-(2+3)",
			},
			want: "-2-(2+3)",
		},
		{
			name: "10",
			args: args{
				expr: "2/(3)",
			},
			want: "2/3",
		},
		{
			name: "11",
			args: args{
				expr: "(2)/(3)",
			},
			want: "2/3",
		},
		{
			name: "12",
			args: args{
				expr: "(2*(3+4)*5)/6",
			},
			want: "2*(3+4)*5/6",
		},
		{
			name: "13",
			args: args{
				expr: "1+(-1)",
			},
			want: "1+-1",
		},
		{
			name: "14",
			args: args{
				expr: "((2*((2+3)-(4*6))+(8+(7*4))))",
			},
			want: "2*(2+3-4*6)+8+7*4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParenthesisRemoval(tt.args.expr); got != tt.want {
				t.Errorf("ParenthesisRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
