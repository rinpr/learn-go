package channel

import "fmt"

func F(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}