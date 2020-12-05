package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	day2inputFile = "day2_input.txt"
)

type password struct {
	min int
	max int
	char uint8
	p string
}

func (p *password) IsValid() bool {
	count := 0
	for i:=0; i<len(p.p); i+=1 {
		if p.p[i] == p.char {
			count += 1
		}
	}
	if count >= p.min && count <= p.max {
		return true
	}
	return false
}

func (p *password) IsValid2() bool {
	a, b := p.p[p.min-1] == p.char, p.p[p.max-1] == p.char
	if (a || b ) && !(a && b) {
		return true
	}
	return false
}

func readPasswords() ([]password, error) {
	curDir, _ := os.Getwd()
	rPath := path.Join(curDir, day2inputFile)
	f, err := os.Open(rPath)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(f)
	passwords := make([]password, 0)
	for {
		l, err := reader.ReadString('\n')
		p := password{}
		parts := strings.Split(l, ": ")
		p.p = parts[1]
		parts = strings.Split(parts[0], " ")
		p.char = parts[1][0]
		parts = strings.Split(parts[0], "-")
		min, _ := strconv.ParseInt([]byte(parts[0]))
		max, _ := strconv.ParseInt([]byte(parts[1]))
		p.min = int(min)
		p.max = int(max)
		passwords = append(passwords, p)
		if err != nil {
			if err == io.EOF {
				return passwords, nil
			}
			return passwords, err
		}
	}
}

func main() {
	passwords, err := readPasswords()
	if err != nil {
		fmt.Printf("%v", err)
	}
	count := 0
	for _, p := range passwords {
		if p.IsValid() {
			count += 1
		}
	}

	fmt.Println(count)

	count = 0
	for _, p := range passwords {
		if p.IsValid2() {
			count += 1
		}
	}

	fmt.Println(count)
}