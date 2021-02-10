package ch03

import "testing"

func TestComma(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "",
			input: "",
			want:  "",
		},
		{
			name:  "300",
			input: "300",
			want:  "300",
		},
		{
			name:  "36,300",
			input: "36300",
			want:  "36,300",
		},
		{
			name:  "1,236,300",
			input: "1236300",
			want:  "1,236,300",
		},
	}
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			if got := Comma(tc.input); got != tc.want {
				t.Fatalf("%s:\n\t- want: %s\n\t-  got: %s",
					tc.name, tc.want, got)
			}
		})
	}
}
