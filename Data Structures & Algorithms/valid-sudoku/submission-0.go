func isValidSudoku(board [][]byte) bool {
	lines := make([][]byte, 9)
	columns := make([][]byte, 0, 9)
	boxes := make(map[string][]byte, 9)
	for row := 0; row < len(board); row++ {
		lines[row] = make([]byte, 0, 9)
		for col := 0; col < len(board[row]); col++ {
			current := board[row][col]

			if row == 0 {
				columns = make([][]byte, 9)
			}

			boxRow := row / 3
			boxCol := col / 3
			trgt := fmt.Sprintf("%d%d", boxRow, boxCol)

			if _, ok := boxes[trgt]; ok {
				boxes[trgt] = append(boxes[trgt], current)
			} else {
				boxes[trgt] = make([]byte, 0, 9)
				boxes[trgt] = append(boxes[trgt], current)
			}

			lines[row] = append(lines[row], current)
			columns[col] = append(columns[col], current)

		}

	}

	validRows := true
	for r := 0; r < len(lines); r++ {
		validRows = validRows && checkLine(lines[r])
	}

	validColumns := true
	for c := 0; c < len(columns); c++ {
		validColumns = validColumns && checkLine(columns[c])
	}
	
	validBoxes := true
	for _, b := range boxes {
		valid := checkLine(b)
		validBoxes = validBoxes && valid
	}

	return validRows && validColumns && validBoxes
}


func checkLine(line []byte) bool {
	localHashMap := make(map[string]bool, 9)
	res := true

	for _, v := range line {
		if v == '.' {
			continue
		}

		if _, exist := localHashMap[string(v)]; exist {
			res = false
		}

		localHashMap[string(v)] = true
	}

	return res
}
