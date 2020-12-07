package main

import (
	"fmt"
	"github.com/tdewolff/parse/strconv"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

const (
	day7inputFile = "day7_input.txt"
)


type bag struct {
	t string
	RawContains []string
	Bags []*rule
}

type rule struct {
	num int
	b *bag
}

func (b *bag) Contains(name string) bool {
	for _, bag := range b.Bags {
		if bag.b.t == name || bag.b.Contains(name) {
			return true
		}
	}
	return false
}

func (b *bag) GetCount() int {
	c := 1
	for _, r := range b.Bags {
		c += r.num * r.b.GetCount()
	}
	return c
}

var bagNodes = map[string]*bag{}

func getRules() error {
	curDir, _ := os.Getwd()
	rPath := path.Join(curDir, day7inputFile)
	data, err := ioutil.ReadFile(rPath)
	if err != nil {
		return err
	}
	splitData := strings.Split(string(data), "\n")
	for _, d := range splitData {
		//light fuchsia bags contain 4 light lavender bags, 5 faded olive bags, 4 plaid cyan bags, 1 striped tomato bag.
		data := strings.Split(d, " bags contain")
		b := &bag{
			t: data[0],
			RawContains: strings.Split(data[1], ","),
		}
		bagNodes[data[0]] = b
	}

	r := regexp.MustCompile(`(\d) (.*) bag[s]?\.?`)
	for _, bagInstance := range bagNodes {
		bagInstance.Bags = make([]*rule, 0)
		for _, c := range bagInstance.RawContains {
			data := r.FindStringSubmatch(c)
			if len(data) > 0 {
				num, _ := strconv.ParseInt([]byte(data[1]))
				ru := rule {
					int(num),
					bagNodes[data[2]],
				}
				bagInstance.Bags = append(bagInstance.Bags, &ru)
			}
		}
	}
	return nil
}

func main() {
	err := getRules()
	if err != nil {
		fmt.Println(err)
		return
	}
	bagName := "shiny gold"
	count := 0
	for _,b := range bagNodes {
		if b.Contains(bagName) {
			count +=1
		}
	}
	fmt.Println(count)

	b := bagNodes[bagName]

	fmt.Println(b.GetCount() - 1)
}