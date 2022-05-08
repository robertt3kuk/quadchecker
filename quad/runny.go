package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	first := atoi(args[0])
	second := atoi(args[1])

	QuadE(first, second)

}

func atoi(s string) int {
	sd := []rune(s)
	res := 0
	for i := 0; i < len(s); i++ {
		res *= 10
		res += int(rune(sd[i] - 48))

	}
	return res

}
