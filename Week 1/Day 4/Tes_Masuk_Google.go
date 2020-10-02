package main

import (
	"fmt"
	"sort"
)

var kumpulanHuruf = []string{"a", "b", "a", "u", "u", "c", "a", "s", "t", "s", "t", "z", "z"}

func main() {
	fmt.Println(hitungVarian(kumpulanHuruf))
}

func hitungVarian(h []string) []map[string]int {
	sort.Strings(h)
	// fmt.Println(h)
	var kumpulan = []map[string]int{} // Slice Map Kosongan
	counter := 1
	for c := 1; c < len(h); c++ {
		if h[c-1] == h[c] {
			counter++
		} else {
			kumpulan = append(kumpulan, map[string]int{h[c-1]: counter})
			counter = 1
		}
		if c == len(h)-1 {
			kumpulan = append(kumpulan, map[string]int{h[c]: counter}) // penutupan, ketika sudah di loop terakhir
		}
	}
	return kumpulan
}
