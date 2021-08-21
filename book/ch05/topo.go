// SPDX-License-Identifier: GPL-2.0
package ch05

import "sort"

func TopoSort(topo map[string][]string) []string {
	var visitAll func([]string)
	var order []string

	seen := make(map[string]bool)
	visitAll = func(items []string) {
		for _, item := range items {
			if _, ok := seen[item]; ok {
				continue
			}
			seen[item] = true
			visitAll(topo[item])
			order = append(order, item)
		}
	}
	var keys []string
	for key := range topo {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
