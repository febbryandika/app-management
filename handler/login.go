package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"simulasi-pair-project/config"
	"strings"
)

func Login() (error, error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan username")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Println("Masukkan username")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	query := "SELECT username, password from users where username= ?"

	db, err := config.GetDB()
	if err != nil {
		return nil, fmt.Errorf("error when connecting to db: %v", err)
	}

	defer db.Close()
	var uname, pwd string
	err = db.QueryRow(query, username).Scan(&uname, &pwd)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("")
		}
	}

	return nil, nil
}
