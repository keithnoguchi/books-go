package ch04

import (
	"io"
	"testing"
)

func TestNewCounter(t *testing.T) {
	tests := []struct {
		name string
		data []rune
		want map[rune]int
	}{
		{
			name: "zero data",
			data: []rune{},
			want: map[rune]int{},
		},
		{
			name: "single rune",
			data: []rune("a"),
			want: map[rune]int{
				'a': 1,
			},
		},
		{
			name: "duplicate runes",
			data: []rune("This is a test"),
			want: map[rune]int{
				'T': 1,
				'h': 1,
				'i': 2,
				's': 3,
				' ': 3,
				'a': 1,
				't': 2,
				'e': 1,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			in, out := io.Pipe()
			go func(out *io.PipeWriter) {
				if _, err := out.Write([]byte(string(tc.data))); err != nil {
					t.Fatalf("%s: write error: %v", tc.name, err)
				}
				out.Close()
			}(out)
			c := NewCounter(in)
			if err := c.Scan(); err != nil {
				t.Fatalf("%s: unexpected start error: %v", tc.name, err)
			}
			for r, want := range tc.want {
				if got := c.Count(r); got != want {
					t.Fatalf("%s: unexpected result for %q:\n\t- want: %d\n\t-  got: %d",
						tc.name, r, want, got)
				}
			}
		})
	}
}
