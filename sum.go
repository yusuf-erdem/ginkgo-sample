package ginkgo_sample

import "fmt"

func Sum(x, y int) int {
	fmt.Println("Received Digits :", x, y)
	return x + y
}

func Sum2(x, y int) int {
	fmt.Println("Received Digits :", x, y)
	return x + y
}

func main() {
	m, n := 3, 4
	fmt.Println(Sum(m, n))

}
