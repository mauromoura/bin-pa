package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4}
	permute(4, values)
}

func binPacking(weights []int, bin int) {

}

func permute(n int, A []int) {
	if n == 1 {
		fmt.Println(A)
	} else {
		for i := 0; i < n-1; i++ {
			permute(n-1, A)
			if n%2 == 0 {
				swap(A, i, n-1)
			} else {
				swap(A, 0, n-1)
			}
		}
		permute(n-1, A)
	}
}

func swap(slice []int, index1 int, index2 int) {
	aux := slice[index1]
	slice[index1] = slice[index2]
	slice[index2] = aux
}
