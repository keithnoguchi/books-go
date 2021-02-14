package v1

import "testing"

func TestAppendInt(t *testing.T) {
	tests := []struct {
		name string
		init []int
		data []int
		want []int
	}{
		{
			name: "append zoro data to zero slice",
			init: []int{},
			data: []int{},
			want: []int{},
		},
		{
			name: "append zero data to slice",
			init: []int{1, 2, 3, 4, 5},
			data: []int{},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "append data to slice",
			init: []int{1, 2, 3, 4, 5},
			data: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := tc.init
			for _, d := range tc.data {
				got = AppendInt(got, d)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("%s: wrong slice size:\n\t- want: %d\n\t-  got: %d",
					tc.name, len(tc.want), len(got))
			}
			for i, want := range tc.want {
				if got[i] != want {
					t.Fatalf("%s: unexpected data\n\t- want: %d\n\t-  got: %d",
						tc.name, want, got[i])
				}
			}
		})
	}
}
