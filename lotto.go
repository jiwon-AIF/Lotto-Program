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
					break
				}
			}
		} else if choice == "Q" {
			os.Exit(1)
		}
	}
}
