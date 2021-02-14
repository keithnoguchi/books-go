package v2

import "testing"

func TestAppendInt(t *testing.T) {
	tests := []struct {
		name string
		init []int
		data []int
		want []int
	}{
		{
			name: "zero data to zero slice",
			init: []int{},
			data: []int{},
			want: []int{},
		},
		{
			name: "data to zero slice",
			init: []int{},
			data: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "data to slice",
			init: []int{1, 2, 3, 4, 5},
			data: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := tc.init
			got = AppendInt(got, tc.data...)
			if len(got) != len(tc.want) {
				t.Fatalf("%s: unexpected length:\n\t- want: %d\n\t-  got: %d",
					tc.name, len(tc.want), len(got))
			}
			for i, want := range tc.want {
				if got[i] != want {
					t.Fatalf("%s: unexpected value:\n\t- want: %d\n\t-  got: %d",
						tc.name, want, got[i])
				}
			}
		})
	}
}
