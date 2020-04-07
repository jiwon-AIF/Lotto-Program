package lotto

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	run()
}

func run() {
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "L" {
		fmt.Println("L")
		fmt.Println(rand.Intn(45) + 1)
	}
}

func RandomInt() int {
	return rand.Intn(45) + 1
}
