package everyday

import "strings"

var board = [][]byte{
	[]byte("abcde"),
	[]byte("fghij"),
	[]byte("klmno"),
	[]byte("pqrst"),
	[]byte("uvwxy"),
	{'z'},
}

var (
	m = 5
	n = 6
)

type position struct {
	x int
	y int
}

func (p1 position) go2Position(p2 position) []byte {
	// 先左右走，再上下走，这样就不会在访问 z 时越界
	// 不过p1 是 z的话要反过来
	if board[p1.x][p1.y] == 'z' {
		return p1.zgo2position(p2)
	}
	y := p2.y - p1.y
	x := p2.x - p1.x
	size := 0
	if x > 0 {
		size += x
	} else {
		size -= x
	}
	if y > 0 {
		size += y
	} else {
		size -= y
	}
	path := make([]byte, 0, size)
	for y > 0 {
		path = append(path, 'R')
		y--
	}
	for y < 0 {
		path = append(path, 'L')
		y++
	}
	for x > 0 {
		path = append(path, 'D')
		x--
	}
	for x < 0 {
		path = append(path, 'U')
		x++
	}
	return path
}

func (p1 position) zgo2position(p2 position) []byte {
	y := p2.y - p1.y
	x := p2.x - p1.x
	size := 0
	if x > 0 {
		size += x
	} else {
		size -= x
	}
	if y > 0 {
		size += y
	} else {
		size -= y
	}
	path := make([]byte, 0, size)
	for x > 0 {
		path = append(path, 'D')
		x--
	}
	for x < 0 {
		path = append(path, 'U')
		x++
	}
	for y > 0 {
		path = append(path, 'R')
		y--
	}
	for y < 0 {
		path = append(path, 'L')
		y++
	}
	return path
}

func alphabetBoardPath(target string) string {
	positions := make(map[byte]position, 26)
	for i, row := range board {
		for j, v := range row {
			positions[v] = position{i, j}
		}
	}
	var builder strings.Builder
	p1 := position{0, 0}
	for i := 0; i < len(target); i++ {
		p2 := positions[target[i]]
		builder.Write(p1.go2Position(p2))
		p1 = p2
		builder.WriteByte('!')
	}
	return builder.String()
}
