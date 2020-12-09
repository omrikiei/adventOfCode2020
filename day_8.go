package main

import (
	"fmt"
	"github.com/tdewolff/parse/strconv"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type operation int
const (
	day8inputFile = "day8_input.txt"
	acc operation = iota
	jmp
	nop
)

type instruction struct {
	Op operation
	value int
	prev *instruction
	next *instruction
}

func (i *instruction) Toggle() {
	if i.Op == jmp {
		i.Op = nop
	} else if i.Op == nop {
		i.Op = jmp
	}
}

func (i instruction) Run(accumulator *int) *instruction {
	next := i.next
	if i.Op == acc {
		*accumulator += i.value
	} else if i.Op == jmp {
		if i.value > 0 {
			for j := 1; j < i.value; j++ {
				next = next.next
			}
		} else {
			for j := 1; j > i.value; j-- {
				next = next.prev
			}
		}
	}
	return next
}

func getOp(o string) operation {
	if o == "jmp" {
		return jmp
	}
	if o == "acc" {
		return acc
	}
	return nop
}

func getInstructions() (*instruction, *instruction, error) {
	curDir, _ := os.Getwd()
	rPath := path.Join(curDir, day8inputFile)
	data, err := ioutil.ReadFile(rPath)
	if err != nil {
		return nil, nil, err
	}
	var instIndex *instruction
	var rootIndex *instruction
	splitData := strings.Split(string(data), "\n")
	for i, data := range splitData {
		d := strings.Split(data, " ")
		op := getOp(d[0])
		value, _ := strconv.ParseInt([]byte(d[1]))
		inst := &instruction{
			Op:    op,
			value: int(value),
			prev:  instIndex,
			next:  nil,
		}
		if instIndex != nil {
			instIndex.next = inst
		}
		instIndex = inst
		if i == 0 {
			rootIndex = inst
		}
	}
	return rootIndex, instIndex, nil
}

func main() {
	fmt.Println("jmp", jmp, "acc", acc, "nop", nop)
	inst, last, _ := getInstructions()
	rootInst := inst
	visited := map[*instruction]bool{}
	accumulator := 5
	for visited[inst] == false {
		fmt.Println(inst.Op, inst.value)
		visited[inst] = true
		inst = inst.Run(&accumulator)
	}
	fmt.Println(accumulator)
	fmt.Println(last)

	changed := map[*instruction]bool{}
	for {
		inst := rootInst
		accumulator = 0
		visited = map[*instruction]bool{}
		var lastChanged *instruction
		for inst != nil && visited[inst] == false {
			visited[inst] = true
			if lastChanged == nil && changed[inst] == false && (inst.Op == jmp || inst.Op == nop){
				lastChanged = inst
				changed[inst] = true
				lastChanged.Toggle()
			}
			inst = inst.Run(&accumulator)
		}
		if inst != nil && lastChanged != nil {
			lastChanged.Toggle()
		}
		if inst == nil {
			break
		}

	}
	fmt.Println(accumulator)


}