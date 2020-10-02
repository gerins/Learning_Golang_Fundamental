package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nama string = "qweasdqweasd123"
	fmt.Println(len(nama))
	fmt.Println(`nama =`, nama, `=>`, reflect.TypeOf(nama))
}
