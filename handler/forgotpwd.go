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

func ForgotPassword() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan username")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Println("Masukkan new password")
	new_password, _ := reader.ReadString('\n')
	new_password = strings.TrimSpace(new_password)

	fmt.Println("Masukkan security answer")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error when connecting to db: ", err)
		return
	}

	defer db.Close()
	var u entity.User
	query := "SELECT security_answer FROM users WHERE username = ?"
	err = db.QueryRow(query, username).Scan(&u.AnswerQuestion)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Error: User not Found")
		} else {
			fmt.Println("Error: ", err)
		}
		return
	}
	fmt.Println("-", u.AnswerQuestion, "-", answer)
	if u.AnswerQuestion != answer {
		fmt.Println("Security Answer doesn't match")
		return
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error encrypting password")
		return
	}

	query = "UPDATE users SET password = ? WHERE username = ?"
	_, err = db.Exec(query, hashedPwd, username)
	if err != nil {
		fmt.Println("Error update pwd: ", err)
	}

	fmt.Println()
	fmt.Println("Password updated succesfully")
	fmt.Println()
}
