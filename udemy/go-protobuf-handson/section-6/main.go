package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"section-6/protogen/basic"
)

func main() {
	u1 := basic.User{
		Id:       1,
		Username: "john_doe",
		IsActive: true,
		Password: []byte("secret"),
	}

	fmt.Printf("Hello %s, please enter your password: ", u1.Username)

	reader := bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	password = strings.TrimSpace(password)

	if password == string(u1.Password) {
		fmt.Println("Password is correct!")
	} else {
		fmt.Println("Incorrect password.")
	}
}
