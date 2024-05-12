package handler

import (
	"fmt"
	"log"
	"simulasi-pair-project/config"
	"simulasi-pair-project/entity"
)

func UbahStok() ([]entity.Product, error) {
	db, err := config.GetDB()
	var recapProduct []entity.Product
	if err != nil {
		fmt.Println("Error connecting to database", err)
		return nil, err
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
	fmt.Println("	   PRODUCT LIST")
	fmt.Println("=========================")
	for rows.Next() {
		var p entity.Product
		if err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {

			return nil, err
		}
		recapProduct = append(recapProduct, p)
	}
	return recapProduct, nil

}
