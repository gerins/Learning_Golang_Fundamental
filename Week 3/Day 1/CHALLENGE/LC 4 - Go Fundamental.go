// [LIVE CODE - 4] Go Fundamental
// 15 Juni 2020
// Nama : Garin Prakoso

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type group struct {
	namaGroup    string
	daftarKontak []string
}

var listKontak []string
var stringNamaGroup []string
var listGroup []group

func main() {
	if len(listGroup) == 0 {
		readFile()
	}
	mainMenu()
	dataLog("Daftar Kontak.txt", listGroup)

	for {
		var inputMenu string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputMenu = scanner.Text()

		switch inputMenu {
		case "1":
			addKontak()
			break
		case "2":
			addGroup()
			break
		case "3":
			viewListKontak()
		case "4":
			ViewListGroup()
		case "5":
			showKontakWithGroupName()
		case "6":
			os.Exit(0)
		default:
			// main()
		}
	}
}

func mainMenu() {
	fmt.Println("----------------------------------")
	fmt.Println("Main Menu")
	fmt.Println("----------------------------------")
	fmt.Println("1. Add Kontak")
	fmt.Println("2. Add Grup")
	fmt.Println("3. View List Kontak")
	fmt.Println("4. View List Grup")
	fmt.Println("5. Show Contact With Grup Name")
	fmt.Println("6. Exit")
	fmt.Println("Masukan menu yang dipilih :")
}

func addGroup() {
	fmt.Println("----------------------------------")
	fmt.Println("Add New Group")
	fmt.Println("----------------------------------")
	if len(stringNamaGroup) >= 3 {
		fmt.Println("Daftar Group Sudah Penuh")
		time.Sleep(time.Millisecond * 1500)
		main()
	}
	scanner := bufio.NewScanner(os.Stdin)
	var nama string

	for {
		fmt.Print("Nama group harus lebih dari 5 Karakter :")
		scanner.Scan()
		nama = scanner.Text()
		if len(nama) > 5 {
			break
		}
	}

	var newStringTemp []string
	newGroupTemp := group{nama, newStringTemp}

	stringNamaGroup = append(stringNamaGroup, nama)
	listGroup = append(listGroup, newGroupTemp)
	fmt.Println(`Group Berhasil ditambahkan!`)

	time.Sleep(time.Millisecond * 1500)
	main()
}

func addKontak() {
	fmt.Println("----------------------------------")
	fmt.Println("Add Kontak")
	fmt.Println("----------------------------------")
	if len(listGroup) == 0 {
		fmt.Println("Tambahkan group terlebih dahulu")
		time.Sleep(time.Millisecond * 1500)
		main()
	}
	if len(listKontak) >= 5 {
		fmt.Println("Daftar Kontak Sudah Penuh")
		time.Sleep(time.Millisecond * 1500)
		main()
	}
	scanner := bufio.NewScanner(os.Stdin)
	var nama string

	for {
		fmt.Print("Nama harus lebih dari 3 karakter & Tidak diisi Angka :")
		scanner.Scan()
		nama = scanner.Text()
		cekAngka := true
		for _, value := range strings.Split(nama, "") {
			var deretAngka = strings.Split("1234567890", "")
			for _, value2 := range deretAngka {
				if value == value2 {
					cekAngka = false
				}
			}
		}

		if len(nama) > 3 && cekAngka {
			break
		}
	}

	listKontak = append(listKontak, nama)
	fmt.Println("Masukan ke dalam Group")
	for index, value := range stringNamaGroup {
		fmt.Println(index, `.`, value)
	}

	var inputGroup int

	for {
		fmt.Print("Masukan nomor index group :")
		fmt.Scan(&inputGroup)
		if inputGroup < len(listGroup) {
			break
		}
	}

	listGroup[inputGroup].daftarKontak = append(listGroup[inputGroup].daftarKontak, nama)

	fmt.Println("Berhasil ditambahkan ke dalam Group")
	time.Sleep(time.Millisecond * 1500)
	main()
}

func viewListKontak() {
	fmt.Println("----------------------------------")
	fmt.Println("View List Kontak")
	fmt.Println("----------------------------------")
	inputView := 0
	fmt.Println("----------------------------------")
	fmt.Printf("%-15s\t%-15s\n", "Nama", "Group")
	fmt.Println("----------------------------------")

	for _, isi := range listGroup {
		for _, value2 := range isi.daftarKontak {
			fmt.Printf("%-15s\t%-15s\n", value2, isi.namaGroup)
		}
	}

	fmt.Println("")
	fmt.Print("Input 1 untuk kembali ke Main menu :")
	fmt.Scan(&inputView)
	if inputView == 1 {
		main()
	}
}

func ViewListGroup() {
	fmt.Println("----------------------------------")
	fmt.Println("View List Group")
	fmt.Println("----------------------------------")
	inputView := 0
	fmt.Println("----------------------------------")
	fmt.Printf("%-15s\t%-15s\n", "Index", "Group")
	fmt.Println("----------------------------------")

	for index, value := range stringNamaGroup {
		fmt.Printf("%-15d\t%-15s\n", index, value)
	}

	fmt.Println("")
	fmt.Print("Input 1 untuk kembali ke Main menu :")
	fmt.Scan(&inputView)
	if inputView == 1 {
		main()
	}
}

func showKontakWithGroupName() {
	fmt.Println("----------------------------------")
	fmt.Println("Show Contact With Grup Name")
	fmt.Println("----------------------------------")
	var inputGroup int
	for index, value := range listGroup {
		fmt.Println(index, value.namaGroup)
	}
	for {
		fmt.Print("Masukan nomor index group :")
		fmt.Scan(&inputGroup)
		if inputGroup < len(listGroup) {
			break
		}
	}

	fmt.Println("-----------------------------------------------------")
	fmt.Printf("%3s\t%-15s\t%-15s\n", "Index.", "Nama", "Group")
	fmt.Println("-----------------------------------------------------")

	for index, isi := range listGroup[inputGroup].daftarKontak {
		fmt.Printf("%-3d\t%-15s\t%-15s\n", index, isi, stringNamaGroup[inputGroup])
	}

	inputView := 0
	fmt.Println("")
	fmt.Print("Input 1 untuk kembali ke Main menu :")
	fmt.Scan(&inputView)
	if inputView == 1 {
		main()
	}
}

func dataLog(filePath string, values []group) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, value := range values {
		fmt.Fprintln(f, value) // print values to f, one per line
	}
	return nil
}

func readFile() {
	file, err := os.Open("Daftar Kontak.txt")
	if err != nil {
		dataLog("Daftar Kontak.txt", listGroup)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		splitData := strings.Split(strings.Trim(strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), "]", ""), "[", ""), "{}"), " ")
		tempNameKontak := splitData[1:]
		listKontak = append(listKontak, tempNameKontak...)
		listGroup = append(listGroup, group{splitData[0], tempNameKontak})
		stringNamaGroup = append(stringNamaGroup, splitData[0])
	}
}
