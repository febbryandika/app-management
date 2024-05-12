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
	QuerySalesRecap :=
		`
    SELECT s.sale_id, p.product_name, s.product_id, s.quantity, s.sale_date
    FROM sales s
    INNER JOIN products p ON s.product_id = p.product_id
    ` // rekapPenjualan query

	rows, err := db.Query(QuerySalesRecap)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var s entity.Sale
		if err = rows.Scan(&s.SaleID, &s.ProductName, &s.ProductID, &s.Quantity, &s.SalesDate); err != nil {
			return nil, err
		}
		recapSales = append(recapSales, s)
	}

	return recapSales, nil
}
