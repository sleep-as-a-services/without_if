package main

import "fmt"

type do func(l, o, r int) string

var (
	m  = []string{"+", "-", "*"}
	n  = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	fs = []do{
		func(l, o, r int) string { return fmt.Sprintf("%v %s %v", l, m[o], n[r]) },
		func(l, o, r int) string { return fmt.Sprintf("%v %s %v", n[l], m[o], r) },
	}
)

func main() {
	fmt.Printf("%s", g(0, 1, 1, 1))
}

func g(f, l, o, r int) string {
	return fs[f](l, o, r)
}
