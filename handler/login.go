package handler

import (

	"bufio"
	"database/sql"
	"fmt"
	"os"
	"simulasi-pair-project/config"
	"simulasi-pair-project/entity"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Login() (string, error) {

	var user entity.User
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan username")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Println("Masukkan password")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	query := "SELECT username, password from users where username= ?"

	db, err := config.GetDB()
	if err != nil {
		return "", fmt.Errorf("error when connecting to db: %v", err)
	}

	defer db.Close()
	err = db.QueryRow(query, username).Scan(&user.UserName, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user or password doesn't match: ")
		} else {
			return "", err
		}
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("user or password doesn't match: ")
	}

	fmt.Println("Login succeed")
	return user.UserName, nil
}
