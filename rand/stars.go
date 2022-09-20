// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func stars(n int) {
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= 2*n-1; j++ {
// 			if j >= n-(i-1) && j <= n+(i-1) {
// 				fmt.Print("#")
// 			} else {
// 				fmt.Print(" ")
// 			}
// 		}
// 		fmt.Println(" ")
// 	}
// }

// func main() {
// 	fmt.Print("Enter how much rows: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	i, _ := strconv.Atoi(input[:len(input)-1])
// 	stars(i)
// }
