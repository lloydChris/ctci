package problem8

//The better way is to set bool flags for the first row and colum if they contain 0
// them use those to track if the whole row/column should be zeroed.
func zero(input [][]int) {
	cols := make(map[int]bool)

	for _, row := range input {
		shouldZero := false
		for j, e := range row {
			if e == 0 {
				shouldZero = true
				cols[j] = true
			}
		}
		if shouldZero {
			for j := range row {
				row[j] = 0
			}
		}
	}

	for col := range cols {
		for i := range input {
			input[i][col] = 0
		}
	}
}
