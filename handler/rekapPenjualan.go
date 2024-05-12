package handler


import (
	"log"
	"simulasi-pair-project/config"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

func RekapPenjualan() {
	db, err := config.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	QuerySalesRecap := "SELECT sale_id, product_id, quantity, sale_date FROM sales"

	rows, err := db.Query(QuerySalesRecap)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// for rows.Next()
}