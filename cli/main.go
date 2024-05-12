package main

import (
	"bufio"
	"fmt"
	"os"
	"simulasi-pair-project/handler"
	"strings"
)

func Report() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=========================")
		fmt.Println("       Fitur Product")
		fmt.Println("=========================")
		fmt.Println()

		fmt.Println("1. Tambah Produk ")
		fmt.Println("2. Ubah Stok Produk")
		fmt.Println("3. Tambah Staff")
		fmt.Println("4. Rekap Penjualan")
		fmt.Println("5. Exit")
		fmt.Println()
		fmt.Print("Silahkan masukkan nomor: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			handler.TambahProduk()
		case "2":
			handler.UbahStok()
		case "3":
			handler.TambahStaff()
		case "4":
			recap, err := handler.RekapPenjualan()
			if err != nil {
				fmt.Println("Error getting the database : ", err)
			}
			for index, value := range recap {
				fmt.Printf("index : %d\n %+v\n", index, value)
			}
		case "5":
			fmt.Println("Exit fitur....")
			return
		default:
			fmt.Println("Input invalid!!!")
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=========================")
		fmt.Println("Selamat datang di Product")
		fmt.Println("=========================")
		fmt.Println()

		fmt.Println("1. Login ")
		fmt.Println("2. Register")
		fmt.Println("3. Forgot Password")
		fmt.Println("4. Exit")
		fmt.Println()
		fmt.Print("Silahkan masukkan nomor: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			user, err := handler.Login()
			if err != nil {
				fmt.Println("Error connecting to db", err)
				break
			}
			fmt.Println("Hello, ", user)
			Report()
		case "2":
			handler.Register()
		case "3":
			handler.ForgotPassword()
		case "4":
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Input invalid:")
		}
	}

}
