package minimaxsum

func MinMaxSum(si []int) (int, int) {

	var min, max, iAux int
	siAux := make([]int, len(si))

	copy(siAux, si)

	aux := si[0]
	for i := 1; i < len(si); i++ {
		if aux > si[i] {
			aux = si[i]
			iAux = i
		}
	}

	si = append(si[:iAux], si[iAux+1:]...)

	//fmt.Println(si)

	for i := 0; i < len(si); i++ {
		max += si[i]
	}

	aux = siAux[0]
	for i := 1; i < len(siAux); i++ {
		if aux < siAux[i] {
			aux = siAux[i]
			iAux = i
		}
	}

	siAux = append(siAux[:iAux], siAux[iAux+1:]...)

	//fmt.Println(siAux)

	for i := 0; i < len(siAux); i++ {
		min += siAux[i]
	}

	return min, max
}
