package defers

import "fmt"

func Add(var1, var2 float64) {
	res := var1 + var2
	fmt.Println("Result:", res)
}

func Loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i =", i)
	}
}

func DeferLoop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println("j =", j)
	}
}