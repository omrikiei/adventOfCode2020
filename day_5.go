package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	day5inputFile = "day5_input.txt"
	maxRow = 128
	maxSeat = 8
)

type seat struct {
	code string
	row int
	seat int
}

func (s *seat) setRow() {
	m, row := maxRow, 0
	for _, c := range s.code[:7] {
		m = m/2
		if c == 'B' {
			row += m
		}
	}
	s.row = row
}

func (s *seat) setSeat() {
	m, se := maxSeat, 0
	for _, c := range s.code[7:] {
		m = m/2
		if c == 'R' {
			se += m
		}
	}
	s.seat = se
}

func (s seat) getSeat() int {
	return s.row * 8 + s.seat
}

func NewSeat(code string) *seat {
	s := &seat{code: code}
	s.setRow()
	s.setSeat()
	return s
}

func getSeats() ([]*seat, error) {
	curDir, _ := os.Getwd()
	rPath := path.Join(curDir, day5inputFile)
	data, err := ioutil.ReadFile(rPath)
	if err != nil {
		return nil, err
	}
	seats := make([]*seat, 0)
	splitData := strings.Split(string(data), "\n")
	for _, d := range splitData {
		seats = append(seats, NewSeat(d))
	}
	return seats, nil
}

func main() {
	m := 0
	seats, _ := getSeats()
	for _, s := range seats {
		t := s.getSeat()
		if t > m {
			m = t
		}
	}
	fmt.Println(m)

	allPossibleSeats := make([]bool, 1024)
	for _, s := range seats {
		allPossibleSeats[s.getSeat()] = true
	}
	for i, s := range allPossibleSeats {
		if i != 0 && i != len(allPossibleSeats) -1 &&
			s == false && allPossibleSeats[i-1] == true && allPossibleSeats[i+1] == true {
			fmt.Println(i)
			break
		}
	}
}