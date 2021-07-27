package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckPrime(number int) bool {
	isPrime := true
	if number <= 3 || number > 100 {
		isPrime = false
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			if number%4 == 1 {
				isPrime = true
			} else {
				isPrime = false
			}
		}
	}
	return isPrime
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	number, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if CheckPrime(int(number)) {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}
