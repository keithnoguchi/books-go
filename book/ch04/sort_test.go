package ch04

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want []int
	}{
		{
			name: "empty data",
			data: []int{},
			want: []int{},
		},
		{
			name: "sorted data",
			data: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "random data",
			data: []int{1, 99, 51, 22, 0},
			want: []int{0, 1, 22, 51, 99},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			Sort(tc.data)
			if len(tc.data) != len(tc.want) {
				t.Fatalf("%s: unexpected length:\n\t- want: %d\n\t-  got: %d",
					tc.name, len(tc.data), len(tc.want))
			}
			for i, got := range tc.data {
				if got != tc.want[i] {
					t.Fatalf("%s: unexpected value in %d:\n\t- want: %d\n\t-  got: %d",
						tc.name, i, tc.want[i], got)
				}
			}
		})
	}
}
