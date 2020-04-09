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
			fmt.Println("행운의 숫자는 :")

			for i := 0; i < len(a); i++ { //숫자 6개를 뽑기 위한 for문
				a[i] = rand.Intn(45) + 1 // 1~45 숫자 중 랜덤으로 하나를 뽑아 a[0]~a[6]에 저장

				for j := 0; j < i; j++ { // 중복제거를 위한 for문 : 현재 a[]에 저장된 랜덤숫자와 이전 a[]에 들어간 숫자 비교
					if a[i] == a[j] {
						i--
					}
				}
			}
			for k := 0; k <= 5; k++ { // 채워진 배열을 출력하기 위한 for문
				fmt.Println(a[k])
				if k == 5 {
					fmt.Println("입니다.")
					break
				}
				var n1, n2, n3, n4, n5, n6 int
				n1 = 1
				n2 = 2
				n3 = 3
				n4 = 4
				n5 = 5
				n6 = 6

				var r1, r2, r3, r4, r5, r6 int
				r1 = a[k]
				r2 = a[k]
				r3 = a[k]
				r4 = a[k]
				r5 = a[k]
				r6 = a[k]
				fmt.Println(n1 == r1)
				fmt.Println(n2 == r2)
				fmt.Println(n3 == r3)
				fmt.Println(n4 == r4)
				fmt.Println(n5 == r5)
				fmt.Println(n6 == r6)
			}
		}
	}
}
