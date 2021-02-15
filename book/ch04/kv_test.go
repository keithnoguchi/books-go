package ch04

import "testing"

func TestKVAdd(t *testing.T) {
	tests := []struct {
		name string
		init KV
		data [][]string
		want KV
	}{
		{
			name: "null strings",
			init: NewKV(),
			data: [][]string {
				[]string {
					"",
					"",
					"",
				},
			},
			want: KV{
				m: map[string]int {
					`["" "" ""]`: 1,
				},
			},
		},
		{
			name: "duplicate strings",
			init: NewKV(),
			data: [][]string {
				[]string {
					"I",
					"am",
					"lucky",
				},
				[]string {
					"I",
					"am",
					"lucky",
				},
			},
			want: KV{
				m: map[string]int {
					`["I" "am" "lucky"]`: 2,
				},
			},
		},
		{
			name: "different strings",
			init: NewKV(),
			data: [][]string {
				[]string {
					"I",
					"am",
					"lucky",
				},
				[]string {
					"I",
					"am",
					"not",
					"lucky",
				},
			},
			want: KV{
				m: map[string]int {
					`["I" "am" "lucky"]`: 1,
					`["I" "am" "not" "lucky"]`: 1,
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := tc.init
			for _, data := range tc.data {
				got.Add(data)
			}
			if !tc.want.Same(&got) {
				t.Fatalf("%s: unexpected result:\n\t- want: %v\n\t-  got: %v",
					tc.name, tc.want, got)
			}
		})
	}
}
