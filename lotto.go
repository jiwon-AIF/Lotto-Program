package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "L" {

			a := make([]int, 6) // int형 배열 선언
			b := make([]int, 6) // int형 배열 선언
			c := make([]int, 6)

			for i := 0; i < len(a); i++ { //숫자 6개를 뽑기 위한 for문
				a[i] = rand.Intn(45) + 1 // 1~45 숫자 중 랜덤으로 하나를 뽑아 a[0]~a[6]에 저장

				for j := 0; j < i; j++ { // 중복제거를 위한 for문 : 현재 a[]에 저장된 랜덤숫자와 이전 a[]에 들어간 숫자 비교
					if a[i] == a[j] {
						i--
					}
				}
			}

			// 당첨번호 설정
			b[0] = 1
			b[1] = 2
			b[2] = 3
			b[3] = 4
			b[4] = 5
			b[5] = 6

			fmt.Println(a)
			fmt.Println(b)

			// fmt.Println("추천 번호: ")
			var cnt = 0
			for k := 0; k < len(a); k++ { // 채워진 배열을 출력하기 위한 for문

				for l := 0; l < len(b); l++ {
					if a[k] == b[l] {
						c[cnt] = a[k]
						cnt++
					}
				}

				// var r1, r2, r3, r4, r5, r6 int
				// r1 = a[k]
				// r2 = a[k]
				// r3 = a[k]
				// r4 = a[k]
				// r5 = a[k]
				// r6 = a[k]

				// if n1 == r1 {
				// 	fmt.Println(r1)
				// } else if n2 == r2 {
				// 	fmt.Println(r2)
				// } else if n3 == r3 {
				// 	fmt.Println(r3)
				// } else if n4 == r4 {
				// 	fmt.Println(r4)
				// } else if n5 == r5 {
				// 	fmt.Println(r5)
				// } else if n6 == r6 {
				// 	fmt.Println(r6)
				// }
			}

			fmt.Println(c)
			for x := 0; x < len(c); x++ {
				if c[x] == 0 {
					fmt.Println(c[:x])
					break
				}
			}

		} else if choice == "Q" {
			os.Exit(1)
		}
	}
}
