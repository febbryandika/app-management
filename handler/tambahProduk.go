package handler

import (
	"bufio"
	"fmt"
	"os"
	"simulasi-pair-project/config"
	"simulasi-pair-project/entity"
	"strconv"
	"strings"
)

func TambahProduk() {
	reader := bufio.NewReader(os.Stdin)
	var produk entity.Product

	fmt.Println("Masukan Nama Produk: ")
	name, _ := reader.ReadString('\n')
	produk.Name = strings.TrimSpace(name)

	fmt.Println("Masukan Harga Produk: ")
	harga, _ := reader.ReadString('\n')
	harga = strings.TrimSpace(harga)
	produk.Price, _ = strconv.ParseFloat(harga, 64)

	fmt.Println("Masukan Stok Produk: ")
	stok, _ := reader.ReadString('\n')
	stok = strings.TrimSpace(stok)
	produk.Stock, _ = strconv.Atoi(stok)

	query := "INSERT INTO products (product_name, price, stock) VALUES (?, ?, ?)"
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error connecting to database", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(query, produk.Name, produk.Price, produk.Stock)
	if err != nil {
		fmt.Println("Error inserting data into database", err)
		return
	}

	fmt.Println("Sucessfully insert new product data")
}
