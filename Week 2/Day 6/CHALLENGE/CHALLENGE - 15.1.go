//[CHALLENGE - 15.1] Bermain Dengan File
// 13 Juni 2020
// Nama : Garin Prakoso

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type daftarMahasiswa struct {
	nama    string
	umur    int
	jurusan string
}

var listMahasiswa = []daftarMahasiswa{} // Slice of struct

func main() {
	if len(listMahasiswa) == 0 {
		readFile()
	}
	mainMenu()

	for {
		var inputMenu string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputMenu = scanner.Text()

		switch inputMenu {
		case "1":
			tambahMahasiswa()
			break
		case "2":
			hapusMahasiwa()
			break
		case "3":
			viewMahasiswa()
		case "4":
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
	fmt.Println("1. Add Mahasiswa")
	fmt.Println("2. Delete Mahasiswa")
	fmt.Println("3. View Mahasiswa")
	fmt.Println("4. Exit")
	fmt.Println("Masukan menu yang dipilih :")
}

func tambahMahasiswa() {
	fmt.Println("----------------------------------")
	fmt.Println("Add Mahasiswa")
	fmt.Println("----------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	var nama, jurusan string
	var umur int

	for {
		fmt.Print("Nama (3-20) Karakter :")
		scanner.Scan()
		nama = scanner.Text()
		if len(nama) > 3 && len(nama) < 20 {
			break
		}
	}

	for {
		fmt.Print("Umur (min 17 tahun) :")
		fmt.Scan(&umur)
		if umur >= 17 {
			break
		}
	}

	for {
		fmt.Print("Jurusan (maks 10 karakter) :")
		fmt.Scan(&jurusan)
		if len(jurusan) <= 10 {
			break
		}
	}

	listMahasiswa = append(listMahasiswa, daftarMahasiswa{nama, umur, jurusan})
	dataLog("dataMahasiswa.txt", listMahasiswa)
	fmt.Println(`Berhasil ditambahkan!`)
	main()
}

func hapusMahasiwa() {
	fmt.Println("----------------------------------")
	fmt.Println("Delete Mahasiswa")
	fmt.Println("----------------------------------")
	if len(listMahasiswa) < 1 {
		fmt.Println(`Daftar mahasiswa kosong`)
		time.Sleep(time.Millisecond * 1500)
		main()
	}
	if len(listMahasiswa) == 1 {
		index := len(listMahasiswa) - 1
		// listMahasiswa = append(listMahasiswa[:index], listMahasiswa[:index]...)
		listMahasiswa = listMahasiswa[:index]
	} else {
		index := len(listMahasiswa) - 1
		// listMahasiswa = append(listMahasiswa[:index], listMahasiswa[:index-1]...)
		listMahasiswa = listMahasiswa[:index]
	}
	fmt.Println(`Mahasiswa yang terakhir berhasil dihapus`)
	dataLog("dataMahasiswa.txt", listMahasiswa)
	time.Sleep(time.Millisecond * 1500)
	main()
}

func byIndex() {
	fmt.Println("----------------------------------")
	fmt.Println("View By Index")
	fmt.Println("----------------------------------")
	var inputIndex int
	for {
		fmt.Print("Masukan index yang ingin ditampilkan :")
		fmt.Scan(&inputIndex)
		if inputIndex < len(listMahasiswa) {
			break
		} else {
			fmt.Println("Mahasiswa dengan index", inputIndex, `tidak dapat ditemukan`)
			fmt.Println("Ingin Kembali ke Main Menu? 1 = Yes | 0 = Kembali Masukan Index")
			fmt.Scan(&inputIndex)
			if inputIndex == 1 {
				main()
			} else if inputIndex == 0 {
				continue
			}
		}
	}
	fmt.Println(" ")
	fmt.Println(`Nama : `, listMahasiswa[inputIndex].nama)
	fmt.Println(`Umur : `, listMahasiswa[inputIndex].umur)
	fmt.Println(`Jurusan : `, listMahasiswa[inputIndex].jurusan)
	fmt.Println(" ")
	fmt.Print("Input 1 untuk kembali ke Main menu :")
	fmt.Scan(&inputIndex)
	if inputIndex == 1 {
		main()
	}
}

func viewAll() {
	fmt.Println("----------------------------------")
	fmt.Println("View All")
	fmt.Println("----------------------------------")
	inputView := 0
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("%3s\t%-15s\t%-15s\t%-15s\n", "Index.", "Nama", "Umur", "Jurusan")
	fmt.Println("-----------------------------------------------------")

	for index, isi := range listMahasiswa {
		fmt.Printf("%-3d\t%-15s\t%-15d\t%-15s\n", index, isi.nama, isi.umur, isi.jurusan)
	}

	fmt.Print("Input 1 untuk kembali ke Main menu :")
	fmt.Scan(&inputView)
	if inputView == 1 {
		main()
	}
}

func viewMahasiswa() {
	fmt.Println("----------------------------------")
	fmt.Println("View Mahasiswa")
	fmt.Println("----------------------------------")
	fmt.Println("1. View By Index")
	fmt.Println("2. View All")
	var inputView int
	for {
		fmt.Print("Masukan menu yang dipilih :")
		fmt.Scan(&inputView)
		if inputView < 3 && inputView > 0 {
			break
		}
	}

	switch inputView {
	case 1:
		byIndex()
		break
	case 2:
		viewAll()
		break
	}
}

func dataLog(filePath string, values []daftarMahasiswa) error {
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
	file, err := os.Open("dataMahasiswa.txt")
	if err != nil {
		dataLog("dataMahasiswa.txt", listMahasiswa)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator
		splitData := strings.Split(strings.Trim(scanner.Text(), "{}"), " ")
		if len(splitData) == 3 {
			umur, _ := strconv.Atoi(splitData[1])
			listMahasiswa = append(listMahasiswa, daftarMahasiswa{splitData[0], umur, splitData[2]})
		} else {
			umur, _ := strconv.Atoi(splitData[2])
			listMahasiswa = append(listMahasiswa, daftarMahasiswa{splitData[0] + " " + splitData[1], umur, splitData[3]})
		}
	}
}
