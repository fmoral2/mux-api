// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"sort"
// 	"time"
// )

// type Adress struct {
// 	adress string
// }

// type User struct {
// 	name     string
// 	lastName string
// 	age      int
// 	ad       *Adress
// }

// func (aa *User) String() string {
// 	return fmt.Sprintf("User{name:%v, lastName:%v, age:%v, Adress: %s}", aa.name, aa.lastName, aa.age, aa.ad)
// }

// func (bb *Adress) String() string {
// 	return fmt.Sprintf("Adress{address:%s}", bb.adress)
// }

// func main() {

// 	t := &Tree{}
// 	t.insert(15)
// 	t.insert(2)
// 	t.insert(10)
// 	t.insert(30)

// 	fmt.Println(t.root.search(2))
// 	fmt.Println(t.root.data, t.root.left.data, t.root.right.data)

// 	sl := []int{3, 55, 2, 39, 87, 4, 8, 8, 2}

// 	fmt.Println("palindrome function: ", IsPalindrome("occo"))

// 	slInt := []int{3, 4, 5, 3, 2, 98, 8}

// 	sort.IntSlice(slInt).Sort()

// 	fmt.Println(slInt)

// 	// convert slice to interface
// 	sli := make([]interface{}, 0)

// 	types := reflect.TypeOf(sli)
// 	fmt.Println(types)

// 	sli = append(sli, []string{"ola", "hello"})
// 	sli = append(sli, slInt)
// 	fmt.Println(sli)

// 	sls := []string{"joao", "pedro", "chico", "oco", "chico", "pedro"}
// 	slStruct := []User{
// 		{name: "joao", lastName: "soares", age: 18},
// 		{name: "maria", lastName: "soares", age: 18},
// 		{name: "ana", lastName: "soares", age: 18},
// 	}

// 	sort.Slice(slStruct, func(i, j int) bool {
// 		return slStruct[i].name < slStruct[j].name

// 	})
// 	fmt.Println(slStruct)

// 	// channels and routines
// 	canal1 := make(chan interface{}, 10)
// 	canal1 <- sls
// 	messageStringSlice := <-canal1

// 	canal2 := make(chan []string, 10)
// 	canal2 <- messageStringSlice.([]string)
// 	strSl := <-canal2

// 	go func() {
// 		// pretend to use list previously created in another thread
// 		sorted := quickSort(sl, 0, len(sl)-1)
// 		fmt.Println("list sorted:", sorted)
// 		intsWithoutDup := removeDup(sorted)
// 		fmt.Println(intsWithoutDup)

// 	}()
// 	time.Sleep(3 * time.Second)
// 	go func() {
// 		// pretending to use in another thread
// 		fmt.Println(strSl)
// 	}()
// 	time.Sleep(3 * time.Second)
// }

// // recursive , choose i pivot (after sorted -> correct position , items to the left are smaller and to the right are larger) and partition the array
// // find partition position por pivot -  increment i until i <= pivot then decrement j until > pivot then swap then if i < j and continue next iteration until all elements
// // to the left are smaller then pivot and to the right are bigger , then recursively quick sort the 2 remaining lists

// // quicksort

// //buble sort

// //merge sort

// //find and remove dups

// //  insert

// // insertNode

// // search

// // is balanced

// // is bst

// //traversals
