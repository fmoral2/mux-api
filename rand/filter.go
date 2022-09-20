//go:build disabled

package main

import (
	"fmt"
	"strings"
)

type Employees struct {
	name string
	role string
	wage int
}

// func main() {
// 	users :=
// 		[]Employees{
// 			{"John Doe", "gardener", 1},
// 			{"Roger Roe", "driver", 2},
// 			{"Paul Smith", "programmer", 3},
// 			{"Lucia Mala", "teacher", 3},
// 			{"Patrick Connor", "shopkeeper", 1},
// 		}

// 	res :=
// 		filterPredicate(users, func(s ...string) bool {
// 			role := "driver"
// 			name := "Connor"

// 			for _, v := range s {
// 				if strings.Contains(v, role) || strings.Contains(v, name) {
// 					return true
// 				}
// 			}
// 			retur false
// 		})

// 	ss := []string{"chico", "joao"}
// 	a := strings.Join(ss, "\n  >>>")
// 	fmt.Println(a)

// 	for _, v := range res {
// 		fmt.Println(v)
// 	}
// }

// // e func = predicate , function that takes a single argument and returns a bool.

// func filterPredicate(e []Employees, f func(...string) bool) []Employees {

// 	a := make([]Employees, 0)

// 	for _, v := range e {
// 		if f(v.role, v.name) {
// 			a = append(a, v)
// 		}
// 	}
// 	return a
// }

func main() {

	users :=
		[]Employees{
			{"John Doe", "gardener", 1},
			{"Roger Roe", "driver", 2},
			{"Paul Smith", "programer", 3},
			{"Lucia Mala", "teacher", 3},
			{"Patrick Connor", "shopkeeper", 1},
		}

	res := predicateFilter(users, func(s ...string) bool {

		d := "john doe"
		c := "programer"

		for _, value := range s {
			if strings.Contains(value, c) || strings.EqualFold(value, d) {
				return true
			}
		}

		return false
	})
	fmt.Println(res)
}

func predicateFilter(e []Employees, f func(...string) bool) []Employees {

	var newList []Employees

	for _, v := range e {
		if f(v.name, v.role) {
			newList = append(newList, v)
		}
	}

	return newList
}
