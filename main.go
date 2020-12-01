package main

import "fmt"

func main() {
	var num, qoldiq, temp int
	var reverse int = 0

	fmt.Print("Enter any positive integer : ")
	fmt.Scan(&num)

	temp = num
	for {
		qoldiq = num % 10
		reverse = reverse*10 + qoldiq
		num /= 10

		if num == 0 {
			break
		}
	}

	if temp == reverse {
		fmt.Printf("%d is a Palindrome", temp)
	} else {
		fmt.Printf("%d is not a Palindrome", temp)
	}
}
