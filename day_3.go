package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	day3inputFile = "day3_input.txt"
	tree = '#'
)

type steps struct {
	right int
	down int
}

func getTrees() ([]string, error) {
	curDir, _ := os.Getwd()
	rPath := path.Join(curDir, day3inputFile)
	data, err := ioutil.ReadFile(rPath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), err
}

func solveDayThree(trees []string, right, down int) int {
	s, count := 0, 0
	for i := 0; i < len(trees); i += down {
		line := trees[i]
		if line[s%len(line)] == tree {
			count += 1
		}
		s += right
	}
	return count
}

func main() {
	trees, err := getTrees()
	if err != nil {
		fmt.Sprintf("%v", err)
	}

	fmt.Println(solveDayThree(trees, 3, 1))
	stepsCheck, res := []steps{{1, 1},{3, 1}, {5, 1},{ 7, 1}, {1, 2}}, 1
	for _, s := range stepsCheck {
		res *= solveDayThree(trees, s.right, s.down)
	}
	fmt.Println(res)
}
