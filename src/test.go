package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	fmt.Printf("%v", s)
}

func addUp(x, y int) (z int) {
	z = x + y
	return
}

func toSum(in ...int) (sum int) {
	for _, i := range in {
		sum += i
	}
	return
}

func useFunc(f func(int, int) int, x, y int) int {
	return f(x, y)
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func feedMe(portion int, eaten int) int {
	eaten += portion
	if eaten >= 5 {
		return eaten
	}
	return feedMe(portion, eaten)
}