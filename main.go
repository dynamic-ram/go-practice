package main

import (
	"fmt"
)

func operation1(num1 int, num2 int) int {
	return num1 + num2
}
func operation2(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}
func main() {
	/*Reading Input and converting String to Int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a Value: ")
	input, _ := reader.ReadString('\n')
	len, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)*/
	var a int = 100
	var b int = 80
	var num = operation1(a, b) // calling function add()
	fmt.Println("SUM:", num)

	var c int = 100
	var d int = 80
	var add, sub = operation2(c, d)
	fmt.Println("SUM:", add)
	fmt.Println("SUB:", sub)
}
