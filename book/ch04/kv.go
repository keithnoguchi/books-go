package ch04

import "fmt"

type KV struct {
	m map[string]int
}

func NewKV() KV { return KV{m: make(map[string]int)} }
func (kv *KV) Add(list []string) string {
	k := kv.Key(list)
	kv.m[k]++
	return k
}
func (kv *KV) Count(list []string) int { return kv.m[kv.Key(list)] }
func (kv *KV) Same(other *KV) bool {
	if len(kv.m) != len(other.m) {
		return false
	}
	for k, v := range kv.m {
		if other.m[k] != v {
			return false
		}
	}
	return true
}
func (kv *KV) Key(list []string) string { return fmt.Sprintf("%q", list) }
