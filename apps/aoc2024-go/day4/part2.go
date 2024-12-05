package day4

func FindCrosses(board [][]rune) int {
	result := 0
	runeSum := rune('m') + rune('s')
	for x, row := range board {
		for y, cell := range row {
			if cell == 'a' {
				if x-1 < 0 || x+1 >= len(board) || y-1 < 0 || y+1 >= len(board[0]) {
					continue
				}
				// Check the upper right and lower left diagonals
				cross1 := board[x-1][y-1] + board[x+1][y+1]
				cross2 := board[x+1][y-1] + board[x-1][y+1]
				if cross1 == runeSum && cross2 == runeSum {
					result += 1
				}
			}
		}
	}
	return result
}

// Failed attempt to implement the solution using a trie:
// type Point struct {
// 	x int
// 	y int
// }

// func newPoint(y, x int) *Point {
// 	return &Point{x: x, y: y}
// }

// func (p Point) Add(q Point) *Point {
// 	return newPoint(p.y+q.y, p.x+q.x)
// }

// func (p Point) Mul(s int) *Point {
// 	return newPoint(p.y*s, p.x*s)
// }

// type diagonalTrie struct {
// 	node *trieNode
// }

// func NewDiagonalTrie() *diagonalTrie {
// 	return &diagonalTrie{node: NewTrieNode()}
// }

// func (t *diagonalTrie) insert(word string) {
// 	node := t.node
// 	for _, c := range word {
// 		if c < 'a' || c > 'z' {
// 			fmt.Printf("Invalid character: %c\n", c)
// 			continue
// 		}
// 		if !node.containsKey(c) {
// 			node.put(c, NewTrieNode())
// 		}
// 		node = node.get(c)
// 	}
// 	node.flag = true
// }

// type resultEntry struct {
// 	endPoint Point
// 	dir      Point
// }

// func newResultEntry(endPoint *Point) *resultEntry {
// 	return &resultEntry{endPoint: *endPoint}
// }

// func (t *diagonalTrie) search(board [][]rune, node *trieNode, i, j int, path []rune, result *map[string][]resultEntry, prevCell []int) {
// 	// If the current cell is out of the board or the current cell is already visited or the current cell is not in the trie, return
// 	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
// 		return
// 	}
// 	c := board[i][j]
// 	if c == '#' || !node.containsKey(c) {
// 		return
// 	}
// 	// If the current cell is in the trie, move to the next cell
// 	node = node.get(c)
// 	path = append(path, c)
// 	if node.flag {
// 		resultEntry := *newResultEntry(newPoint(i, j))
// 		// Update the result entry with the direction
// 		resultEntry.dir = *newPoint(i-prevCell[0], j-prevCell[1])
// 		(*result)[string(path)] = append((*result)[string(path)], resultEntry)
// 	}
// 	// Mark the current cell as visited
// 	board[i][j] = '#'
// 	// Move to the next cell.
// 	// If the previous cell is nil, move to all 8 directions
// 	if prevCell == nil {
// 		// Go clockwise: DownLeft, UpLeft, UpRight, DownRight
// 		t.search(board, node, i+1, j-1, path, result, []int{i, j}) // DownLeft
// 		t.search(board, node, i-1, j-1, path, result, []int{i, j}) // UpLeft
// 		t.search(board, node, i-1, j+1, path, result, []int{i, j}) // UpRight
// 		t.search(board, node, i+1, j+1, path, result, []int{i, j}) // DownRight
// 	}
// 	// If the previous cell is not nil, move to the next cell in the direction of the previous cell
// 	if prevCell != nil {
// 		// Go clockwise: DownLeft, UpLeft, UpRight, DownRight
// 		if i > prevCell[0] && j < prevCell[1] {
// 			t.search(board, node, i+1, j-1, path, result, []int{i, j}) // DownLeft
// 		}
// 		if i < prevCell[0] && j < prevCell[1] {
// 			t.search(board, node, i-1, j-1, path, result, []int{i, j}) // UpLeft
// 		}
// 		if i < prevCell[0] && j > prevCell[1] {
// 			t.search(board, node, i-1, j+1, path, result, []int{i, j}) // UpRight
// 		}
// 		if i > prevCell[0] && j > prevCell[1] {
// 			t.search(board, node, i+1, j+1, path, result, []int{i, j}) // DownRight
// 		}
// 	}
// 	// Re-set the current cell to the original value
// 	board[i][j] = c
// }

// func FindDiagonalWords(board [][]rune, words []string) map[string][]resultEntry {
// 	result := make(map[string][]resultEntry)
// 	t := NewDiagonalTrie()
// 	for _, word := range words {
// 		t.insert(word)
// 	}
// 	for i, row := range board {
// 		for j, _ := range row {
// 			t.search(board, t.node, i, j, []rune{}, &result, nil)
// 		}
// 	}
// 	return result
// }

// func FindCrosses(board [][]rune, words []string) map[string]int {
// 	diagonalMatches := FindDiagonalWords(board, words)
// 	result := make(map[string]int)
// 	countedEndPoints := make(map[Point]bool)
// 	for word, entries := range diagonalMatches {
// 		for _, entry := range entries {
// 			// Calculate the cross end points. We don't need the current direction and we dont need the exact opposite.
// 			// So, calculate the cross end points for the two possible crossing directions
// 			crossEndPoint1 := entry.endPoint.Add(*newPoint(entry.dir.y, entry.dir.x*-1).Mul(len(word)))
// 			shouldTryCrossEndpoint1 := crossEndPoint1.x >= 0 && crossEndPoint1.x <= len(board[0]) && crossEndPoint1.y >= 0 && crossEndPoint1.y <= len(board[0])
// 			crossEndPoint2 := entry.endPoint.Add(*newPoint(entry.dir.y*-1, entry.dir.x).Mul(len(word)))
// 			shouldTryCrossEndpoint2 := crossEndPoint2.x >= 0 && crossEndPoint2.x <= len(board[0]) && crossEndPoint2.y >= 0 && crossEndPoint2.y <= len(board[0])
// 			for _, candidate := range entries {
// 				// Skip this entry
// 				if candidate.endPoint.x == entry.endPoint.x && candidate.endPoint.y == entry.endPoint.y {
// 					continue
// 				}
// 				// If an entry has already been counted, skip it
// 				if countedEndPoints[candidate.endPoint] {
// 					continue
// 				}
// 				if shouldTryCrossEndpoint1 && candidate.endPoint.x == crossEndPoint1.x && candidate.endPoint.y == crossEndPoint1.y {
// 					result[word]++
// 					countedEndPoints[candidate.endPoint] = true
// 				}
// 				if shouldTryCrossEndpoint2 && candidate.endPoint.x == crossEndPoint2.x && candidate.endPoint.y == crossEndPoint2.y {
// 					result[word]++
// 					countedEndPoints[candidate.endPoint] = true
// 				}
// 			}
// 		}
// 	}
// 	return result
// }
