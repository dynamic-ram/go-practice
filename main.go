package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func funval(num int) {
	num = 5
}

func funptr(ptr *int) {
	*ptr = 5
}

func main() {
	//Reading Input and converting String to Int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a Value: ")
	input, _ := reader.ReadString('\n')
	len, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	fmt.Println(len)
	i := 1
	funval(i)
	fmt.Println("funval:", i)
	fmt.Println("pointer:", &i)
	funptr(&i)
	fmt.Println("funptr:", i)
	fmt.Println("pointer:", &i)
}
