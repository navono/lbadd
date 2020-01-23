package scanner

import "testing"

func Benchmark_trie(b *testing.B) {
	trie := NewRuneTrie()
	for k, v := range keywordsWithA {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithB {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithC {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithD {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithE {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithF {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithG {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithH {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithI {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithJ {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithK {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithL {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithM {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithN {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithO {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithP {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithQ {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithR {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithS {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithT {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithU {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithV {
		trie.Put(k, v)
	}
	for k, v := range keywordsWithW {
		trie.Put(k, v)
	}
	benchMarkGetTrie("INITIALk", b, trie)
}

func benchMarkGetTrie(val string, b *testing.B, trie *RuneTrie) {
	for n := 0; n < b.N; n++ {
		trie.Get("SELECT")
	}
}
