/////////////////////////////////////////////////////////////////////////
// FOR LOOP

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	for i := 1; i <= 15; i++ {

// 		defer fmt.Println(i) // 10 (1+2+3+4)
// 	}

// 	fmt.Println("Finish")

// }

/////////////////////////////////////////////////////////////////////////
// CONVERT TYPE DATA

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {

// 	var nilaiAngka = strconv.Itoa(2020) // int ke string
// 	fmt.Print(nilaiAngka, ` adalah `)
// 	fmt.Printf("%T\n", nilaiAngka)

// 	var nilaiInt, _ = strconv.Atoi("1234") // string ke int
// 	fmt.Print(nilaiInt, ` adalah `)
// 	fmt.Printf("%T\n", nilaiInt)

// 	// strconv.ParseBool
// }

///////////////////////////////////////////////////////////////////////////
// SWITCH CASE

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	Saturday := 7
// 	today := 6
// 	switch Saturday {
// 	case today + 0:
// 		fmt.Println("Today.")
// 	case today + 1:
// 		fmt.Println("Tomorrow.")
// 	case today + 2:
// 		fmt.Println("In two days.")
// 	default:
// 		fmt.Println("Too far away.")
// 	}
// }

///////////////////////////////////////////////////////////////////////////
// BUFIO

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan nama anda: ")
	name, _ := reader.ReadString('\n')
	// name, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println(name)

}
