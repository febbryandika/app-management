package main

import (
	"fmt"
	"simulasi-pair-project/config"
)

func main() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

}
