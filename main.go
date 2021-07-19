package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a Value less than 10: ")
	input, _ := reader.ReadString('\n')
	value, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	for i := value; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is Even")
		} else {
			fmt.Println(i, "is Odd")
		}
		switch i {
		case 1:
			fmt.Println(i, "is January")
		case 7:
			fmt.Println(i, "is July")
		case 8:
			fmt.Println(i, "is August")
		case 9:
			fmt.Println(i, "is September")
		case 10:
			fmt.Println(i, "is October")
		default:
			fmt.Printf("%v is not in the list\n", i)
		}
	}
	fmt.Println("\nHello World")
}
