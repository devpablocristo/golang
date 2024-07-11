package diagonaldifference

import "fmt"

func DiagonalDifference(arr [][]int32) int32 {
	var l int
	for _, lr := range arr {
		l = len(lr)
		break
	}
	fmt.Println("l:", l)

	var r1 int32
	for i := 0; i < l; i++ {
		r1 += arr[i][i]
	}
	fmt.Println("r1:", r1)

	var r2, j int32
	for i := l - 1; i >= 0; i-- {
		r2 += arr[i][j]
		j++
	}
	fmt.Println("r2:", r2)

	r := r1 - r2

	if r < 0 {
		r *= int32(-1)
	}

	fmt.Println("r:", r)
	return r

}
