package handler


import (
	"bufio"
	"fmt"
	"os"
	"simulasi-pair-project/config"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Register() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Masukkan password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Print("Masukkan Security Question: ")
	question, _ := reader.ReadString('\n')
	question = strings.TrimSpace(question)

	fmt.Print("Masukkan Security Answer: ")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error when generating pwd: %v", err)
		return
	}

	query := "INSERT INTO users (username, password, security_question, security_answer) VALUES (?,?,?,?)"
	db, err := config.GetDB()
	if err != nil {
		fmt.Printf("Error when connecting to db : %v", err)
		return
	}

	defer db.Close()

	_, err = db.Exec(query, username, hashedPwd, question, answer)
	if err != nil {
		fmt.Printf("Error while registering data: %v", err)
	}

	fmt.Println("User Registered Succesfully")


}
