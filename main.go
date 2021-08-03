package main

import "fmt"

/*
func operation1(num1 int, num2 int) int {
	return num1 + num2
}
func operation2(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}
*/

//Narcissitic Function
/*
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

func operation(nums ...int) int {
	var total = 0
	for _, v := range nums {
		total += v
	}
	return total
}
*/

/* Closure Function
func operation(num1 int) func() int {
	i := 0
	return func() int {
		i += 1
		return num1 + i
	}
}

//Recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
//Defer
func c() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}
func a() {
	fmt.Println("Fresco")
}

func b() {
	fmt.Println("Play")
}
*/
//Panic and Recover Function
func panicFunction() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("PANIC!!")
}

func main() {
	/*Reading Input and converting String to Int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a Value: ")
	input, _ := reader.ReadString('\n')
	len, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	//Variatic function loop
	l := int(len)
	var x = make([]int, l)
	for i := 0; i < l; i++ {
		input, _ := reader.ReadString('\n')
		y, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
		x[i] = int(y)
	}
	//Narcissitic conditions
	if isNarcissistic(int(i)) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	var a int = 100
	var b int = 80
	var num = operation1(a, b) // calling function add()
	fmt.Println("SUM:", num)

	var c int = 100
	var d int = 80
	var add, sub = operation2(c, d)
	fmt.Println("SUM:", add)
	fmt.Println("SUB:", sub)
	fmt.Println("SUM:", operation(x...))

	//Function Closure
	var a int = 100
	var num = operation(a)
	fmt.Println("num=", num())
	fmt.Println("num=", num())
	fmt.Println("num=", num())
	var num1 = operation(a)
	fmt.Println("num1=", num1())
	fmt.Println("num1=", num1())

	//Recursion
	var a int = 4
	var num = factorial(a)
	fmt.Println(num)

	a()
	defer b()
	c()
	panic("Panic Occured!")
	//fmt.Println("This line is not displayed")
	*/
	//Panic and Recover
	panicFunction()

	fmt.Println("This line appears when recover is used inside deferred function")

}
