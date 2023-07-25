package main

import "fmt"

func main() {
	count := 1
	for count < 6 {
		count += count
	}
	fmt.Println(count) // 8
}
