package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var productList []map[string]string
var kategoriProduk []string

func clearScreen() {
	osRunning := runtime.GOOS
	if osRunning == "linux" || osRunning == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if osRunning == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	mainMenu()
}

func mainMenu() {
	fmt.Println("**************************************")
	fmt.Println("Aplikasi")
	fmt.Println("**************************************")
	fmt.Println("1. Buat Produk Baru")
	fmt.Println("2. Daftar Produk")
	fmt.Println("4. Keluar")
	fmt.Println("Pilihan menu dari 1-4")
}

func listProductForm() {
	fmt.Println("Halaman Daftar Produk")
	fmt.Println("**************************************")
	fmt.Printf("%3s\t%-20s\t%-20s\t%-20s\n", "No.", "Kode Produk", "Nama Produk", "Kategori")
	for idx, product := range productList {
		fmt.Printf("%-3d\t%-20s\t%-20s\t%-20s\n", idx, product["productCode"], product["productName"], kategoriProduk[idx])
	}
}

func newProductForm() {
	var productCode string
	var productName string
	var saveProductConfirmation string
	var saveKategori string

	fmt.Println("Halaman Buat Produk Baru")
	fmt.Println("**************************************")
	fmt.Print("Product Code : ")
	fmt.Scanln(&productCode)
	fmt.Print("Product Name : ")
	fmt.Scanln(&productName)
	fmt.Print("Kategori Product : ")
	fmt.Scanln(&saveKategori)
	fmt.Printf("Produk %s kategori %s dengan kode %s akan disimpan (y/n) ? ", productName, saveKategori, productCode)
	fmt.Scanln(&saveProductConfirmation)

	if saveProductConfirmation == "y" {
		newProduct := make(map[string]string)
		newProduct["productCode"] = productCode
		newProduct["productName"] = productName
		productList = append(productList, newProduct)
		kategoriProduk = append(kategoriProduk, saveKategori)
		clearScreen()
	}
}

func main() {
	mainMenu()
	for {
		var selectedMenu string

		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			newProductForm()
			break
		case "2":
			listProductForm()
			break
		case "4":
			os.Exit(0)
		default:
			clearScreen()
		}
	}
}
