package main

import "fmt"

func main() {
	var mahasiswa = []map[string]string{
		map[string]string{"nama": "Didik Prabowo", "alamat": "Wonogiri"},
		map[string]string{"nama": "Rudi Subgyo", "alamat": "Solo"},
	}

	mahasiswa = append(mahasiswa,
		map[string]string{"nama": "Gerin Prakoso", "alamat": "Bekasi"},
		map[string]string{"nama": "Azzalia Kurnianingrum", "alamat": "Madiun"},
		map[string]string{"nama": "Kucing Hitam", "alamat": "Jakarta"},
	)

	for i, v := range mahasiswa {
		fmt.Println(i, "= nama :", v["nama"], "alamat:", v["alamat"])
	}
}
