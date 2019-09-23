package project7

//Requires a square matrix of nxn
func rotate(image [][]byte) {
	n := len(image)
	var x int = n / 2
	var y int = n - 1

	for i := 0; i < x; i++ {
		for j := i; j < y-i; j++ {
			temp := image[i][j]
			image[i][j] = image[y-j][i]
			image[y-j][i] = image[y-i][y-j]
			image[y-i][y-j] = image[j][y-i]
			image[j][y-i] = temp
		}
	}
}
