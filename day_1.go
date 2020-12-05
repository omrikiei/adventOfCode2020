package main

import (
	"bufio"
	"fmt"
	"github.com/tdewolff/parse/strconv"
	"io"
	"os"
	"path"
	"sort"
)

const (
	day1inputFile = "day1_input.txt"
	year = 2020
)

func readInputs() ([]int, error) {
	curDir, _ := os.Getwd()
	rPath := path.Join(curDir, day1inputFile)
	f, err := os.Open(rPath)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(f)
	nums := make([]int, 0)
	for {
		l, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nums, nil
			}
			return nums, err
		}
		i, _ := strconv.ParseInt(l)
		if err != nil {
			return nums, err
		}
		nums = append(nums, int(i))
	}
}

func getFirst(nums []int, target int) int {
	sort.Ints(nums)
	s, e := 0, len(nums) - 1
	for {
		y := nums[s] + nums[e]
		if y == target {
			fmt.Println(nums[s]*nums[e])
			return nums[s]*nums[e]
		}
		if y > target {
			e -= 1
		} else {
			s += 1
		}
		if s > e {
			return -1
		}
	}
}

func getSecond(nums []int) int {
	sort.Ints(nums)
	fmt.Println(len(nums))
	a, b, c := 0,1,2
	for nums[a] + nums[b] + nums[c] < year {
		c += 1
	}
	for {
		t := year - nums[c]
		fmt.Println(c, t)
		r := getFirst(nums[:c], t)
		if r != -1 {
			return r * nums[c]
		}
		c -= 1
		if c == 0 {
			return -1
		}
	}
}

func main() {
	nums, _ := readInputs()
	fmt.Println(getFirst(nums, year))
	fmt.Println(getSecond(nums))
}
