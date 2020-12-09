package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	day9inputFile = "day9_input.txt"
)

func getNumbersFromInput() []int64 {
	curDir, _ := os.Getwd()
	rPath := path.Join(curDir, day9inputFile)
	data, err := ioutil.ReadFile(rPath)
	if err != nil {
		return []int64{}
	}
	splitData := strings.Split(string(data), "\n")
	nums := make([]int64, len(splitData))
	for i, data := range splitData {
		n, _ := strconv.ParseInt(data, 10, 64)
		nums[i] = n
	}
	return nums
}

func isValid(num int64, prev []int64) bool {
	for _, n1 := range prev {
		k := num - n1
		if k < 0 {
			continue
		}
		for _, n2 := range prev {
			if n2 == n1 { continue }
			if k - n2 == 0 {
				return true
			}
		}
	}
	return false
}

func getSequence(num int64, nums[]int64) []int64 {
	for i, _ := range nums {
		sum := int64(0)
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == num {
				return nums[i:j]
			}
			if sum > num {
				break
			}
		}
	}
	return []int64{}
}

func main() {
	nums := getNumbersFromInput()
	n := int64(0)
	for i := 25; i < len(nums); i++ {
		if !isValid(nums[i], nums[i-25:i]) {
			fmt.Println(i, nums[i])
			n = nums[i]
			break
		}
	}

	seq := getSequence(n, nums)
	min, max := n, int64(0)
	for _, num := range seq {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	fmt.Println("min", min, "max", max)
	fmt.Println(min+max)

}