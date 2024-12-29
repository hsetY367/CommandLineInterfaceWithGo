package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bxcodec/faker/v3"
)

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i))*127 + 128),
		int(math.Cos(f*float64(i))*127 + 128),
		int(math.Sin(f*float64(i))*127 + 128)
}
func main() {
	var phrases []string

	for i := 1; i < 20; i++ {
		phrases = append(phrases, faker.Sentence())
	}

	output := strings.Join(phrases[:], "; ")

	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
}
