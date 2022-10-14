package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	imprimir()

}

//min max sumatoria
func miniMaxSum(arr []int32) {
	// Write your code here

}

func imprimir() {
	l := 5
	for i := 0; i <= l; i++ {
		fmt.Printf("%"+strconv.Itoa(l)+"s\n", strings.Repeat("#", i))
	}
}

func diagonalDifference(arr [][]int32) int32 {
	var sumDiagPpal int32 = 0
	var sumDiagSec int32 = 0
	for i := range arr {
		sumDiagPpal += arr[i][i]
		sumDiagSec += arr[i][len(arr)-i-1]
	}
	return int32(math.Abs(float64(sumDiagPpal - sumDiagSec)))

}
