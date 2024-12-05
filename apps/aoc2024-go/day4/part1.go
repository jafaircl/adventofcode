package day4

import "fmt"

type trieNode struct {
	flag  bool
	links []*trieNode
}

func NewTrieNode() *trieNode {
	return &trieNode{flag: false, links: make([]*trieNode, 26)}
}

func (n *trieNode) get(c rune) *trieNode {
	return n.links[c-'a']
}

func (n *trieNode) put(c rune, node *trieNode) {
	n.links[c-'a'] = node
}

func (n *trieNode) containsKey(c rune) bool {
	return n.links[c-'a'] != nil
}

type trie struct {
	node *trieNode
}

func NewTrie() *trie {
	return &trie{node: NewTrieNode()}
}

func (t *trie) insert(word string) {
	node := t.node
	for _, c := range word {
		if c < 'a' || c > 'z' {
			fmt.Printf("Invalid character: %c\n", c)
			continue
		}
		if !node.containsKey(c) {
			node.put(c, NewTrieNode())
		}
		node = node.get(c)
	}
	node.flag = true
}

func (t *trie) search(board [][]rune, node *trieNode, i, j int, path []rune, result *map[string]int, prevCell []int) {
	// If the current cell is out of the board or the current cell is already visited or the current cell is not in the trie, return
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	c := board[i][j]
	if c == '#' || !node.containsKey(c) {
		return
	}
	// If the current cell is in the trie, move to the next cell
	node = node.get(c)
	path = append(path, c)
	if node.flag {
		(*result)[string(path)] = (*result)[string(path)] + 1
	}
	// Mark the current cell as visited
	board[i][j] = '#'
	// Move to the next cell.
	// If the previous cell is nil, move to all 8 directions
	if prevCell == nil {
		// Go clockwise: Down, DownLeft, Left, UpLeft, Up, UpRight, Right, DownRight
		t.search(board, node, i+1, j, path, result, []int{i, j})   // Down
		t.search(board, node, i+1, j-1, path, result, []int{i, j}) // DownLeft
		t.search(board, node, i, j-1, path, result, []int{i, j})   // Left
		t.search(board, node, i-1, j-1, path, result, []int{i, j}) // UpLeft
		t.search(board, node, i-1, j, path, result, []int{i, j})   // Up
		t.search(board, node, i-1, j+1, path, result, []int{i, j}) // UpRight
		t.search(board, node, i, j+1, path, result, []int{i, j})   // Right
		t.search(board, node, i+1, j+1, path, result, []int{i, j}) // DownRight
	}
	// If the previous cell is not nil, move to the next cell in the direction of the previous cell
	if prevCell != nil {
		// Go clockwise: Down, DownLeft, Left, UpLeft, Up, UpRight, Right, DownRight
		if i > prevCell[0] && j == prevCell[1] {
			t.search(board, node, i+1, j, path, result, []int{i, j}) // Down
		} else if i > prevCell[0] && j < prevCell[1] {
			t.search(board, node, i+1, j-1, path, result, []int{i, j}) // DownLeft
		} else if i == prevCell[0] && j < prevCell[1] {
			t.search(board, node, i, j-1, path, result, []int{i, j}) // Left
		} else if i < prevCell[0] && j < prevCell[1] {
			t.search(board, node, i-1, j-1, path, result, []int{i, j}) // UpLeft
		} else if i < prevCell[0] && j == prevCell[1] {
			t.search(board, node, i-1, j, path, result, []int{i, j}) // Up
		} else if i < prevCell[0] && j > prevCell[1] {
			t.search(board, node, i-1, j+1, path, result, []int{i, j}) // UpRight
		} else if i == prevCell[0] && j > prevCell[1] {
			t.search(board, node, i, j+1, path, result, []int{i, j}) // Right
		} else if i > prevCell[0] && j > prevCell[1] {
			t.search(board, node, i+1, j+1, path, result, []int{i, j}) // DownRight
		}
	}
	// Re-set the current cell to the original value
	board[i][j] = c
}

func FindWords(board [][]rune, words []string) map[string]int {
	result := make(map[string]int)
	t := NewTrie()
	for _, word := range words {
		t.insert(word)
	}
	for i, row := range board {
		for j, _ := range row {
			t.search(board, t.node, i, j, []rune{}, &result, nil)
		}
	}
	return result
}
