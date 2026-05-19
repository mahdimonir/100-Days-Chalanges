package main

func main() {
	a := 10
	b := a
	c := &a // c is a pointer to a

	a = 20

	d := 30
	e := &d
	*e = 40

	println("Value of a:", a)
	println("Value of b:", b)
	println("Value of c:", *c)
	println("Value of d:", d)
	println("Value of e:", *e)

	pointer()
}

// run: `go run main.go pointer.go`