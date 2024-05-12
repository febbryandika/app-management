package handler

import (
	"fmt"
	"log"
	"simulasi-pair-project/config"
	"simulasi-pair-project/entity"
)

func RekapPenjualan() ([]entity.Sale, error) {
	db, err := config.GetDB()
	var recapSales []entity.Sale
	if err != nil {
		return recapSales, fmt.Errorf("error getting the database : %v", err)
	}
	QuerySalesRecap := "SELECT sale_id, product_id, quantity, sale_date FROM sales" // rekapPenjualan query

	rows, err := db.Query(QuerySalesRecap)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var s entity.Sale
		if err = rows.Scan(&s.ID, &s.ProductID, &s.Quantity, &s.SalesDate); err != nil {
			return nil, err
		}
		recapSales = append(recapSales, s)
	}

	return recapSales, nil
}
