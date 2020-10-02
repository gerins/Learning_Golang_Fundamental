package main

import (
	"fmt"
)

func main() {
	var alamatIntLangsung = new(int) // bikin langsung berupa alamat kosongan
	fmt.Println(alamatIntLangsung)   // masih berupa alamat
	fmt.Println(*alamatIntLangsung)  // isi dari alamat tersebut = masih kosongan
	
	*alamatIntLangsung = 2 * 2
	fmt.Println(*alamatIntLangsung)
}
