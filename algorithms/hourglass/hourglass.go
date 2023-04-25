package main

type Hourglass struct {
	Hg [3][3]int
}

type Hourglass2 struct {
	Hg [][]int32
}

func main() {

}

func HighestHourglassSum(a [6][6]int) int {
	var sum, r int
	hg := GetHourglass(a)
	for i := 0; i < 16; i++ {
		sum = SumHourglass(hg[i].Hg)
		if r < sum {
			r = sum
		}
	}
	return r
}

func SumHourglass(h [3][3]int) int {

	sum1 := h[0][0] + h[0][1] + h[0][2]
	sum2 := h[1][1]
	sum3 := h[2][0] + h[2][1] + h[2][2]

	sumTotal := sum1 + sum2 + sum3

	return sumTotal
}

func GetHourglass(a [6][6]int) [16]Hourglass {
	var hs [16]Hourglass

	for hc := 0; hc < 16; hc++ {
		for c := 0; c < 3; c++ {
			for r := 0; r < 3; r++ {
				switch {
				case hc == 0:
					hs[hc].Hg[r][c] = a[r][c]
				case hc == 1:
					hs[hc].Hg[r][c] = a[r+1][c]
				case hc == 2:
					hs[hc].Hg[r][c] = a[r+2][c]
				case hc == 3:
					hs[hc].Hg[r][c] = a[r+3][c]
				case hc == 4:
					hs[hc].Hg[r][c] = a[r][c+1]
				case hc == 5:
					hs[hc].Hg[r][c] = a[r+1][c+1]
				case hc == 6:
					hs[hc].Hg[r][c] = a[r+2][c+1]
				case hc == 7:
					hs[hc].Hg[r][c] = a[r+3][c+1]
				case hc == 8:
					hs[hc].Hg[r][c] = a[r][c+2]
				case hc == 9:
					hs[hc].Hg[r][c] = a[r+1][c+2]
				case hc == 10:
					hs[hc].Hg[r][c] = a[r+2][c+2]
				case hc == 11:
					hs[hc].Hg[r][c] = a[r+3][c+2]
				case hc == 12:
					hs[hc].Hg[r][c] = a[r][c+3]
				case hc == 13:
					hs[hc].Hg[r][c] = a[r+1][c+3]
				case hc == 14:
					hs[hc].Hg[r][c] = a[r+2][c+3]
				case hc == 15:
					hs[hc].Hg[r][c] = a[r+3][c+3]
				}
			}
		}
	}

	return hs
}

func SumHourglass2(h [][]int32) int32 {

	sum1 := h[0][0] + h[0][1] + h[0][2]
	sum2 := h[1][1]
	sum3 := h[2][0] + h[2][1] + h[2][2]

	sumTotal := sum1 + sum2 + sum3

	return sumTotal
}

func GetHourglass2(a [][]int32) []Hourglass2 {
	hs := make([]Hourglass2, 16)

	for hc := 0; hc < 16; hc++ {
		hs[hc] = Hourglass2{
			Hg: make([][]int32, 16),
		}
		for c := 0; c < 3; c++ {
			for r := 0; r < 3; r++ {
				switch {
				case hc == 0:
					hs[hc].Hg[r][c] = a[r][c]
				case hc == 1:
					hs[hc].Hg[r][c] = a[r+1][c]
				case hc == 2:
					hs[hc].Hg[r][c] = a[r+2][c]
				case hc == 3:
					hs[hc].Hg[r][c] = a[r+3][c]
				case hc == 4:
					hs[hc].Hg[r][c] = a[r][c+1]
				case hc == 5:
					hs[hc].Hg[r][c] = a[r+1][c+1]
				case hc == 6:
					hs[hc].Hg[r][c] = a[r+2][c+1]
				case hc == 7:
					hs[hc].Hg[r][c] = a[r+3][c+1]
				case hc == 8:
					hs[hc].Hg[r][c] = a[r][c+2]
				case hc == 9:
					hs[hc].Hg[r][c] = a[r+1][c+2]
				case hc == 10:
					hs[hc].Hg[r][c] = a[r+2][c+2]
				case hc == 11:
					hs[hc].Hg[r][c] = a[r+3][c+2]
				case hc == 12:
					hs[hc].Hg[r][c] = a[r][c+3]
				case hc == 13:
					hs[hc].Hg[r][c] = a[r+1][c+3]
				case hc == 14:
					hs[hc].Hg[r][c] = a[r+2][c+3]
				case hc == 15:
					hs[hc].Hg[r][c] = a[r+3][c+3]
				}
			}
		}
	}

	return hs
}
