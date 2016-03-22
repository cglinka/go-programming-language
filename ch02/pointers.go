package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(*p) // 1
	fmt.Println(p)  // address
	*p = 2
	fmt.Println(x)  // 2
	fmt.Println(*p) // 2
	fmt.Println(p)  // address

	var a, b int
	fmt.Println(&a == &a, &b == &a, &a == nil) // true false false

	var g = f()
	fmt.Println(g)          // address
	fmt.Println(*g)         // 1
	fmt.Println(f() == f()) // false

	v := 1
	incr(&v)       // v now = 2
	fmt.Println(v) // 2
	// fmt.Println(*v)    // invalid indirect of v (type int)
	fmt.Println(&v)       // address
	fmt.Println(incr(&v)) // 3

}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
