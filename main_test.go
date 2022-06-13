package main

import (
	"fmt"
	_ "image/png"
	"strconv"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	s := "m 0 1 22 3"
	l := strings.Split(s, " ")
	x1, _ := strconv.Atoi(l[1])
	y1, _ := strconv.Atoi(l[2])
	x2, _ := strconv.Atoi(l[3])
	y2, _ := strconv.Atoi(l[4])
	fmt.Println(x1, x2, y1, y2)
}
