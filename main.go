package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
    password := "test"
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Hash:", string(hash))

    err = bcrypt.CompareHashAndPassword(hash, []byte(password))
    if err != nil {
        fmt.Println("Password does not match")
    } else {
        fmt.Println("Password matches")
    }
}
