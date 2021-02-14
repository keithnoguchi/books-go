package ch04

import "testing"

func TestStackPush(t *testing.T) {
	tests := []struct {
		name string
		init Stack
		data []int
		want Stack
	}{
		{
			name: "empty data to empty stack",
			init: Stack{},
			data: []int{},
			want: Stack{},
		},
		{
			name: "empty data to stack",
			init: Stack{
				data: []int{1, 2, 3, 4, 5},
			},
			data: []int{},
			want: Stack{
				data: []int{1, 2, 3, 4, 5},
			},
		},
		{
			name: "data to stack",
			init: Stack{
				data: []int{1, 2, 3, 4, 5},
			},
			data: []int{5, 4, 3, 2, 1},
			want: Stack{
				data: []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := tc.init
			for _, x := range tc.data {
				got.Push(x)
			}
			if len(got.data) != len(tc.want.data) {
				t.Fatalf("%s: unexpected length:\n\t- want: %d\n\t-  got: %d",
					tc.name, len(tc.want.data), len(got.data))
			}
			for i, want := range tc.want.data {
				if got.data[i] != want {
					t.Fatalf("%s: unexpected value:\n\t- want: %d\n\t-  got: %d",
						tc.name, want, got.data[i])
				}
			}
		})
	}
}

func TestStackPop(t *testing.T) {
	tests := []struct {
		name string
		data Stack
		want []int
	}{
		{
			name: "zero pop from empty stack",
			data: Stack{},
			want: []int{},
		},
		{
			name: "pop one from one stack",
			data: Stack{
				data: []int{1},
			},
			want: []int{1},
		},
		{
			name: "pop two from two stack",
			data: Stack{
				data: []int{1, 2},
			},
			want: []int{2, 1},
		},
		{
			name: "pop two from three stack",
			data: Stack{
				data: []int{1, 2, 3},
			},
			want: []int{3, 2},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			for i := 0; i < len(tc.want); i++ {
				x, err := tc.data.Pop()
				if err != nil {
					t.Fatalf("%s: unexpected error: %v",
						tc.name, err)
				}
				got = append(got, x)
			}
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
