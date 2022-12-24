package main

import (
	"fmt"

	// https://qiita.com/k-yoshigai/items/cb094cf5ba0a26013f56
	calc "github.com/MatsuoTakuro/workspace-calc"
)

func main() {
	nums := []float64{1.1, 2.2, 3.3}
	fmt.Printf("result: %v\n", calc.Sum(nums))
}
