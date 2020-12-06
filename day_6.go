package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	day6inputFile = "day6_input.txt"
)

type group struct {
	lines []string
}

func (g group) count1() int {
	m := map[int32]int{}
	for _, line := range g.lines {
		for _, c := range line {
			m[c] += 1
		}
	}
	println(m)

	count := 0
	for _, _ = range m {
		count+=1
	}
	return count
}

func (g group) count2() int {
	m := map[int32]int{}
	for _, line := range g.lines {
		for _, c := range line {
			m[c] += 1
		}
	}
	println(m)

	count := 0
	for _, c := range m {
		if c == len(g.lines) {
			count += 1
		}
	}
	return count
}

func getGroups() ([]*group, error) {
	curDir, _ := os.Getwd()
	rPath := path.Join(curDir, day6inputFile)
	data, err := ioutil.ReadFile(rPath)
	if err != nil {
		return nil, err
	}
	groups := make([]*group, 0)
	splitData := strings.Split(string(data), "\n\n")
	for _, d := range splitData {
		groups = append(groups, &group{
			lines: strings.Split(d, "\n"),
		})
	}
	return groups, nil
}

func main() {
	groups, _ := getGroups()
	count := 0
	for _, g := range groups {
		count += g.count1()
	}
	fmt.Println(count)
	count = 0
	for _, g := range groups {
		count += g.count2()
	}
	fmt.Println(count)
}