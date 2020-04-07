package lotto

import (
	"fmt"
	"math/rand"
	"time"
)

func ExampleRandomInt() {
	rand.Seed(time.Now().UnixNano())
	a := make([]int, 6)

	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(45) + 1

		for j := 0; j < i; j++ {
			if a[i] == a[j] {
				i--
			}
		}
	}
	for k := 0; k <= 5; k++ {
		fmt.Println(a[k])
		// Output:
		//
	}
}
