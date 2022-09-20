// package main

// import (
// 	"fmt"

// 	"golang.org/x/crypto/bcrypt"
// )

// const (
// 	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
// 	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
// 	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
// )

// func main() {
// 	password := []byte("CHICO")

// 	// Hashing the password with the default cost of 10
// 	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(hashedPassword))

// 	// Comparing the password with the hash
// 	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
// 	fmt.Println(err) // nil means it is a match

// 	fmt.Println(string(password))
// 	fmt.Println(hashedPassword)

// }
