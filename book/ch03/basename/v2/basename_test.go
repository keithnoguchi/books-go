package basename

import "testing"

func TestBasename(t *testing.T) {
	tests := []struct {
		name string
		input string
		want string
	}{
		{
			name: "a/b/c.go",
			input: "a/b/c.go",
			want: "c",
		},
		{
			name: "c.d.go",
			input: "c.d.go",
			want: "c.d",
		},
		{
			name: "abc",
			input: "abc",
			want: "abc",
		},
	}
	for _, tc := range tests {
		tc := tc
		if got := Basename(tc.input); got != tc.want {
			t.Fatalf("%s:\n\t- want: %v\n\t-  got: %v",
				tc.name, tc.want, got)
		}
	}
}
