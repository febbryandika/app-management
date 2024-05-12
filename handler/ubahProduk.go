package handler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"simulasi-pair-project/config"
	"simulasi-pair-project/entity"
	"strconv"
	"strings"
)

func UbahStok() {
	db, err := config.GetDB()
	var recapProduct []entity.Product
	var produk entity.Product

	if err != nil {
		fmt.Println("Error connecting to database", err)
		return
	}
	defer db.Close()

	QueryToDisplay :=
		`
	SELECT product_id, product_name, price, stock
	FROM products
	`
	rows, err := db.Query(QueryToDisplay)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("=========================")
	fmt.Println("PRODUCT LIST")
	fmt.Println("=========================")
	for rows.Next() {
		var p entity.Product
		if err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {
			return
		}
		recapProduct = append(recapProduct, p)
	}
	for _, product := range recapProduct {
		fmt.Printf("ID : %d | Product Name : %s | Price : %.2f | Stock %d\n", product.ID, product.Name, product.Price, product.Stock)
	}

	QueryToUpdate :=
		`UPDATE products 
	 SET stock = ? 
	 WHERE product_id = ?
	`

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan ID Produk : ")
	ID, _ := reader.ReadString('\n')
	ID = strings.TrimSpace(ID)
	produk.ID, _ = strconv.Atoi(ID)

	fmt.Println("Ubah Stok Produk: ")
	stok, _ := reader.ReadString('\n')
	stok = strings.TrimSpace(stok)
	produk.Stock, _ = strconv.Atoi(stok)

	result, err := db.Exec(QueryToUpdate, produk.Stock, produk.ID)
	if err != nil {
		log.Fatal("Error executing SQL query:", err)
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Error getting updated rows", err)
	}

	fmt.Println("=========================")
	fmt.Printf("UPDATED PRODUCT STOCK FOR PRODUCT ID %d\n", produk.ID)
	fmt.Println("=========================")

	// Retrieve the updated product list
	rows, err = db.Query(QueryToDisplay)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p entity.Product
		if err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {
			return
		}
		recapProduct = append(recapProduct, p)
		fmt.Printf("ID : %d | Product Name : %s | Price : %.2f | Stock %d\n", p.ID, p.Name, p.Price, p.Stock)
	}

	fmt.Println("=========================")
	fmt.Printf("Stok barang diubah menjadi: %d\n", produk.Stock)
	fmt.Printf("Jumlah produk yang diubah: %d\n", rowsUpdated)

}
