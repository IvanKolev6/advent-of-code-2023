package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func firstAndLastDigit(input string) int {
    re := regexp.MustCompile(`\d`)
    digits := re.FindAllString(input, -1)
    if len(digits) == 0 {
        return 0
    }

	concat := ""

    if len(digits) == 1 {
		concat = digits[0] + digits[0]
       
    } else {
		concat = digits[0] +  digits[len(digits)-1]
    }

	result, err :=  strconv.Atoi(concat)
	if err != nil {
        return 0
    }
	return result
}

func main() {
	file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

	if err := scanner.Err(); err != nil {
        panic(err)
    }
	
	sum := 0

	for _, line := range lines {
		digit := firstAndLastDigit(line)
		fmt.Println(digit)
        sum += digit
    }
	
	fmt.Println(sum)
}
