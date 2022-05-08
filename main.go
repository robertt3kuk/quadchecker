package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"

	"github.com/01-edu/z01"
)

func main() {
	var inputr [][]rune
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		inputr = append(inputr, []rune(line))
	}
	lenX, lenY, err := lenght(inputr)
	if err != "" {
		for _, c := range err {
			z01.PrintRune(c)
		}
		return

	}
	input := stringer(inputr)

	quada, qua := QuadA(lenX, lenY)
	quadb, qub := QuadB(lenX, lenY)
	quadc, quc := QuadC(lenX, lenY)
	quadd, qud := QuadD(lenX, lenY)
	quade, que := QuadE(lenX, lenY)
	var ini int

	if reflect.DeepEqual(input, quada) {
		fmt.Print(qua)
		ini++
	}
	if inii(ini) {
		fmt.Printf(" | ")
		ini--
	}
	if reflect.DeepEqual(input, quadb) {
		fmt.Print(qub)
		ini++
	}
	if inii(ini) {
		fmt.Printf(" | ")
		ini--
	}
	if reflect.DeepEqual(input, quadc) {
		fmt.Print(quc)

		ini++
	}
	if inii(ini) {
		fmt.Printf(" | ")
		ini--
	}
	if reflect.DeepEqual(input, quadd) {
		fmt.Print(qud)
		ini++
	}
	if inii(ini) {
		fmt.Printf(" | ")
		ini--
	}
	if reflect.DeepEqual(input, quade) {
		fmt.Print(que)
		ini++

	}

}
func inii(ini int) bool {
	if ini == 1 {
		return true
	}
	return false
}

func lenght(s [][]rune) (x, y int, err string) {
	y = len(s)
	x = len(s[0])
	for _, c := range s {
		if x != len(c) {
			return 0, 0, "Not a Raid function"
		}

	}
	return x, y, ""

}

func QuadA(x, y int) ([]string, string) {
	c1 := 111 // o
	c2 := 111 // o
	yD := 124 // |
	xD := 45  // -
	abc := 1
	return XxY(x, y, c1, c2, yD, xD, abc), "[quadA]"
}

func QuadB(x, y int) ([]string, string) {
	c1 := 47 // /
	c2 := 92 // \
	yD := 42 // *
	xD := 42 // *
	abc := 2
	return XxY(x, y, c1, c2, yD, xD, abc), "[quadB]"
}

func QuadC(x, y int) ([]string, string) {
	c1 := 65 // A
	c2 := 67 // C
	yD := 66 // B
	xD := 66 // B
	abc := 3
	return XxY(x, y, c1, c2, yD, xD, abc), "[quadC]"
}

func QuadD(x, y int) ([]string, string) {
	c1 := 65 // A
	c2 := 67 // C
	yD := 66 // B
	xD := 66 // B
	abc := 1
	return XxY(x, y, c1, c2, yD, xD, abc), "[quadD]"
}

func QuadE(x, y int) ([]string, string) {
	c1 := 65 // A
	c2 := 67 // C
	yD := 66 // B
	xD := 66 // B
	abc := 2
	return XxY(x, y, c1, c2, yD, xD, abc), "[quadE]"
}

func XxY(x, y, c1, c2, yD, xD, abc int) []string {

	delX := x - 2
	delY := y - 2
	var sd []string
	switch abc {
	case 1:
		sd = append(sd, xaxys(x, delX, c1, c2, xD))
		for i := 0; i < delY; i++ {
			sd = append(sd, yaxys(y, delX, x, yD))
		}
		if y > 1 {
			sd = append(sd, xaxys(x, delX, c1, c2, xD))
		}
	case 2:
		sd = append(sd, xaxys(x, delX, c1, c2, xD))
		for i := 0; i < delY; i++ {
			sd = append(sd, yaxys(y, delX, x, yD))
		}
		if y > 1 {
			sd = append(sd, xaxys(x, delX, c2, c1, xD))
		}
	case 3:
		sd = append(sd, xaxys(x, delX, c1, c1, xD))
		for i := 0; i < delY; i++ {
			sd = append(sd, yaxys(y, delX, x, yD))
		}
		if y > 1 {
			sd = append(sd, xaxys(x, delX, c2, c2, xD))
		}

	}
	return sd
}

func xaxys(x, delX, c1, c2, xD int) string {
	var xax string
	xax += string(rune(c1))
	for delixx := 0; delixx < delX; delixx++ {

		xax += string(rune(xD))
	}
	if x > 1 {

		xax += string(rune(c2))

	}
	return xax
}

func yaxys(y, delX, x, yD int) string {
	var yax string

	yax += string(rune(yD))
	for delixx := 0; delixx < delX; delixx++ {

		yax += string(rune(' '))
	}
	if x > 1 {

		yax += string(rune(yD))
	}
	return yax
}
func printif(s []string) {
	for _, c := range s {
		fmt.Printf("%v,\n", c)
	}
}
func stringer(sd [][]rune) []string {
	biqol := make([]string, len(sd))
	for i := 0; i < len(sd); i++ {
		biqol[i] = string(sd[i])

	}
	return biqol
}
