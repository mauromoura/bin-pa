package main

import (
	"fmt"
	"math"
)

func main() {
	//weights := []int{7, 7, 3, 4}
	//bin := 7
	weights := []int{5, 7, 5, 2, 4, 2, 5, 1, 6}
	bin := 10
	recursive := binPackingRecursive(weights, bin)
	firstFit := firstFitBinPacking(weights, bin)
	fmt.Println(recursive, firstFit)
}

func firstFitBinPacking(weights []int, bin int) int {
	binSlice := []int{bin}
	for i := 0; i < len(weights); i++ {
		settled := false
		j := 0
		for !settled && j < len(binSlice) {
			if weights[i] <= binSlice[j] {
				binSlice[j] -= weights[i]
				settled = true
			} else {
				j++
			}
		}
		if !settled {
			binSlice = append(binSlice, bin-weights[i])
		}
	}
	return len(binSlice)
}

func binPackingRecursive(weights []int, bin int) int {
	min := math.MaxInt32
	permute(len(weights), weights, &min, bin)
	return min
}

func bruteForce(A []int, bin int, min *int) {
	count := 1
	binAux := bin
	for i := 0; i < len(A); i++ {
		if A[i] <= binAux {
			binAux -= A[i]
		} else {
			binAux = bin
			binAux -= A[i]
			count++
		}
	}
	if count < *min {
		*min = count
		//fmt.Println(A, count)
	}
}

func permute(n int, A []int, min *int, bin int) {
	if n == 1 {
		bruteForce(A, bin, min)
	} else {
		for i := 0; i < n-1; i++ {
			permute(n-1, A, min, bin)
			if n%2 == 0 {
				swap(A, i, n-1)
			} else {
				swap(A, 0, n-1)
			}
		}
		permute(n-1, A, min, bin)
	}
}

func swap(slice []int, index1 int, index2 int) {
	aux := slice[index1]
	slice[index1] = slice[index2]
	slice[index2] = aux
}
