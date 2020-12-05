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
	day4inputFile = "day4_input.txt"
)

type passport map[string]string
var mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func (p passport) isValid() bool {
	for _, field := range mandatoryFields {
		if _, ok := p[field]; ok == false {
			return false
		}
	}
	return true
}

func (p passport) isValid2() bool {
	if !p.isValid() {
		return false
	}
	if y, _ := strconv.ParseInt([]byte(p["byr"])); y < 1920 || y > 2002 {
		return false
	}

	if y, _ := strconv.ParseInt([]byte(p["iyr"])); y < 2010 || y > 2020 {
		return false
	}

	if y, _ := strconv.ParseInt([]byte(p["eyr"])); y < 2020 || y > 2030 {
		return false
	}

	hgt, _ := p["hgt"]

	if strings.HasSuffix(hgt, "cm") {
		n, _ := strconv.ParseInt([]byte(strings.Split(hgt, "cm")[0]))
		if n < 150 || n > 193 {
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		n, _ := strconv.ParseInt([]byte(strings.Split(hgt, "in")[0]))
		if n < 59 || n > 76 {
			return false
		}
	} else {
		return false
	}

	m, err := regexp.MatchString("^#[0-9a-f]{6}$",p["hcl"])
	if err != nil || !m {
		return false
	}

	validEyeColors := map[string]bool {
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	e, ok := p["ecl"]
	if !ok { return false}
	_, ok = validEyeColors[e]
	if !ok { return false}

	m, err = regexp.MatchString("^[0-9]{9}$", p["pid"])
	if err != nil || !m {
		return false
	}

	return true
}

func NewPassport(passportData string) passport {
	passportData = strings.ReplaceAll(passportData, "\n", " ")
	data := strings.Split(passportData, " ")
	pass := passport{}
	for _,line := range data {
		d := strings.Split(line, ":")
		pass[d[0]] = d[1]
	}
	return pass
}

func getPassports() ([]passport, error) {
	curDir, _ := os.Getwd()
	rPath := path.Join(curDir, day4inputFile)
	data, err := ioutil.ReadFile(rPath)
	if err != nil {
		return nil, err
	}
	passports := make([]passport, 0)
	splitData := strings.Split(string(data), "\n\n")
	for _, d := range splitData {
		passports = append(passports, NewPassport(d))
	}
	return passports, nil
}

func main() {
	passports, _ := getPassports()
	count := 0

	for _, p := range passports {
		if p.isValid() {
			count += 1
		}
	}
	fmt.Println(count)

	count = 0

	for _, p := range passports {
		if p.isValid2() {
			count += 1
		}
	}

	fmt.Println(count)
}