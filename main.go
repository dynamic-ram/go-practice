package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Reading Input and converting String to Int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a Value: ")
	input, _ := reader.ReadString('\n')
	len, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	l := int(len)
	var x = make([]int, l)
	var z int
	var k int
	for i := 0; i < l; i++ {
		input, _ := reader.ReadString('\n')
		y, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
		x[i] = int(y)
		if i == 0 || i%2 == 0 {
			z += x[i]
		} else {
			k += x[i]
		}
	}
	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(k)
	fmt.Println(z - k)
	/* Arrays and Slices
	var x = make([]int, 5, 10)
	x = append(x, 3, 4)
	var y = [5]int{1, 2, 3, 4, 5}
	var z = x[0:3]
	var q = copy(x, z)
	fmt.Println(value)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(q)
	var x = make([]int, 3) //[0,0,0]
	var y = []int{1, 2, 3} //[1,2,3]
	fmt.Println(x)
	copy(x, y) //[1,2,3]
	fmt.Println(y)
	fmt.Println(x)
	x := [6]int{1, 2, 3, 4, 5, 6}
	z := x[3:5]
	fmt.Println(z)
	*/
	//Maps
	/*
		var x = make(map[string]int)
		x["key0"] = 1
		x["key1"] = 9
		fmt.Println(x["key0"])
		delete(x, "key0")
		fmt.Println(x["key0"])
	*/
}
