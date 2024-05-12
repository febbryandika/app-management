package handler

import (
	"bufio"
	"fmt"
	"os"
	"simulasi-pair-project/config"
	"simulasi-pair-project/entity"
	"strings"
)

func TambahStaff() {
	reader := bufio.NewReader(os.Stdin)
	var staff entity.Staff

	fmt.Println("Masukan Nama Staff: ")
	nama, _ := reader.ReadString('\n')
	staff.Name = strings.TrimSpace(nama)

	fmt.Println("Masukan Email Staff: ")
	email, _ := reader.ReadString('\n')
	staff.Email = strings.TrimSpace(email)

	fmt.Println("Masukan Posisi Staff: ")
	posisi, _ := reader.ReadString('\n')
	staff.Position = strings.TrimSpace(posisi)

	query := "INSERT INTO staff (staff_name, email, position) VALUES (?, ?, ?)"
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error connecting to database", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(query, staff.Name, staff.Email, staff.Position)
	if err != nil {
		fmt.Println("Error inserting data into database", err)
		return
	}

	fmt.Println("Sucessfully insert new staff data")
}
