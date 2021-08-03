package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
func operation1(num1 int, num2 int) int {
	return num1 + num2
}
func operation2(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}
*/

func isNarcissistic(number int) bool {
	tmp := number
	digits := len(strconv.Itoa(number))
	sum := 0
	div := 0
	for {
		if tmp <= 0 {
			break
		}
		div = tmp % 10
		temp := 1
		for i := 0; i < digits; i++ {
			temp = temp * div
		}
		sum += temp
		tmp = tmp / 10
	}
	return number == sum
}

func main() {
	//Reading Input and converting String to Int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a Value: ")
	input, _ := reader.ReadString('\n')
	i, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if isNarcissistic(int(i)) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	/*var a int = 100
	var b int = 80
	var num = operation1(a, b) // calling function add()
	fmt.Println("SUM:", num)

	var c int = 100
	var d int = 80
	var add, sub = operation2(c, d)
	fmt.Println("SUM:", add)
	fmt.Println("SUB:", sub)
	*/

}
