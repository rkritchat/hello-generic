package main

import (
	"fmt"
	"strings"
)

func main() {
	r := Filter([]string{"AXA", "X", "Y"}, func(v string) bool { return len(v) > 2 })
	fmt.Println(r)

	x := Filter([]int{1, 5, 7, 12, 8}, func(v int) bool { return v%2 == 0 })
	fmt.Println(x)

	z := IsIn([]int{1, 2, 3, 4, 5, 6, 7, 100}, 100, func(x, y int) bool { return x == y })
	fmt.Println(z)

	target, predicate := "C", func(x, y string) bool {
		return strings.ToLower(x) == strings.ToLower(y)
	}
	y := IsIn([]string{"A", "B", "C"}, target, predicate)
	fmt.Println(y)

	b := bird{}
	Run(b)
	d := dog{}
	Run(d)
}

func Filter[T any](collection []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func IsIn[T any](collection []T, target T, predicate func(T, T) bool) bool {
	for _, item := range collection {
		if predicate(item, target) {
			return true
		}
	}
	return false
}

type animal interface {
	Move()
}

type bird struct{}

func (b bird) Move() { fmt.Println("fly") }

type dog struct{}

func (d dog) Move() { fmt.Println("walk") }

func Run[T animal](a T) {
	a.Move()
}
