package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("////-Welcome To Our Calculator-////")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter No. 1: ")
	n1, _ := reader.ReadString('\n')
	n2, _ := strconv.ParseFloat(strings.TrimSpace(n1), 64)

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Operation: ")
	n3, _ := reader1.ReadString('\n')
	n4 := strings.TrimSpace(n3)

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter No.2: ")
	n5, _ := reader2.ReadString('\n')
	n6, _ := strconv.ParseFloat(strings.TrimSpace(n5), 64)

	if n4 == "+" {

		res := n2 + n6

		fmt.Println("Result is: ", res)

	} else {
		fmt.Println("Error Enetr correct Opteration")
	}
	if n4 == "-" {

		res := n2 - n6

		fmt.Println("Result is: ", res)

	}

	if n4 == "*" {

		res := n2 * n6

		fmt.Println("Result is: ", res)

	}
	if n4 == "/" {

		res := n2 / n6

		fmt.Println("Result is: ", res)

	}

}
