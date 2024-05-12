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
	// DB Connect
	db, err := config.GetDB()
	var recapProduct []entity.Product
	var produk entity.Product

	// Check error connecting DB
	if err != nil {
		fmt.Println("Error connecting to database", err)
		return
	}
	defer db.Close()

	// Displaying the current DB content for products
	QueryToDisplay :=
		`
	SELECT product_id, product_name, price, stock
	FROM products
	`

	// Executing the query
	rows, err := db.Query(QueryToDisplay)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print the content to CLI
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
	// Print the content to CLI using entity as the structure
	for _, product := range recapProduct {
		fmt.Printf("ID : %d | Product Name : %s | Price : %.2f | Stock %d\n", product.ID, product.Name, product.Price, product.Stock)
	}

	// Query for updating the product stock
	QueryToUpdate :=
		`UPDATE products 
	 SET stock = ? 
	 WHERE product_id = ?
	`

	// For reading the input of the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan ID Produk : ")
	ID, _ := reader.ReadString('\n')
	ID = strings.TrimSpace(ID)
	produk.ID, _ = strconv.Atoi(ID)

	fmt.Println("Ubah Stok Produk: ")
	stok, _ := reader.ReadString('\n')
	stok = strings.TrimSpace(stok)
	produk.Stock, _ = strconv.Atoi(stok)

	// Get the current stock of the product
	var currentStock int
	err = db.QueryRow("SELECT stock FROM products WHERE product_id = ?", produk.ID).Scan(&currentStock)
	if err != nil {
		log.Fatal("Error querying current stock:", err)
	}

	// Update the stock of the product
	_, err = db.Exec(QueryToUpdate, produk.Stock, produk.ID)
	if err != nil {
		log.Fatal("Error updating stock:", err)
	}

	fmt.Println("=========================")
	fmt.Println("PRODUCT THAT'S BEEN UPDATED")
	fmt.Println("=========================")

	fmt.Printf("Product ID : %d | Stock (Before) : %d | Stock (After) : %d\n", produk.ID, currentStock, produk.Stock)

	// Retrieve and display the updated stock of the product
	err = db.QueryRow("SELECT product_name, stock FROM products WHERE product_id = ?", produk.ID).Scan(&produk.Name, &produk.Stock)
	if err != nil {
		log.Fatal("Error querying updated stock:", err)
	}

	fmt.Println("=========================")
	fmt.Printf("UPDATED PRODUCT STOCK\n")
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

}
