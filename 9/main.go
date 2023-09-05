package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const cord_size = 10

type cord struct {
	x int
	y int
}

func (c *cord) EuclidianDist(other cord) float64 {
	return math.Sqrt(math.Abs(math.Pow(float64(other.x-c.x), 2) + math.Pow(float64(other.y-c.y), 2)))
}

func (c *cord) Update(move string) {
	if strings.Contains(move, "R") {
		c.x = c.x + 1
	} else if strings.Contains(move, "L") {
		c.x = c.x - 1
	}

	if strings.Contains(move, "U") {
		c.y = c.y + 1
	} else if strings.Contains(move, "D") {
		c.y = c.y - 1
	}
}

func PrintMap(knots []cord) {
	maxX, maxY, minX, minY := 20, 20, -20, -20

	s := ""

	for i := maxY; i >= minY; i-- {
		for j := minX; j <= maxX; j++ {
			found := false
			for w := 0; w < cord_size; w++ {
				if j == knots[w].x && i == knots[w].y {
					s += strconv.Itoa(w)
					found = true
					break
				}
			}
			if !found {
				if j == 0 && i == 0 {
					s += "s"
				} else {
					s += "."
				}
			}
		}
		s += "\n"
	}
	fmt.Println(s)
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scan := bufio.NewScanner(file)

	var visited = map[cord]struct{}{}
	knots := make([]cord, cord_size)

	for scan.Scan() {
		line := scan.Text()
		tokens := strings.Split(line, " ")

		move := tokens[0]
		reps, _ := strconv.Atoi(tokens[1])

		for i := 0; i < reps; i++ {
			knots[0].Update(move) // move head

			for j := 1; j < cord_size; j++ {
				if int(knots[j].EuclidianDist(knots[j-1])) > 1 {
					tailMove := ""
					if knots[j-1].y > knots[j].y {
						tailMove += "U"
					}
					if knots[j-1].y < knots[j].y {
						tailMove += "D"
					}
					if knots[j-1].x > knots[j].x {
						tailMove += "R"
					}
					if knots[j-1].x < knots[j].x {
						tailMove += "L"
					}
					knots[j].Update(tailMove)
				}
			}
			visited[knots[cord_size-1]] = struct{}{}
		}
		// fmt.Println(line)
		// fmt.Println(visited)
		// fmt.Println(knots)
		// PrintMap(knots)
	}
	fmt.Println(len(visited))
}
