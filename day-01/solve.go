package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func createTheDigit(digits []string)int {
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

func firstAndLastDigit(input string) int {
    re := regexp.MustCompile(`\d`)
    digits := re.FindAllString(input, -1)
    if len(digits) == 0 {
        return 0
    }

    the_digit := createTheDigit(digits)

    return the_digit
}

func firstAndLastDigitString(input string) int {
    digitNames := map[string]int{
        "zero":  0,
        "one":   1,
        "two":   2,
        "three": 3,
        "four":  4,
        "five":  5,
        "six":   6,
        "seven": 7,
        "eight": 8,
        "nine":  9,
    }

    var digits []string
    for pos := range input {
        str := input[pos:]
        for key, value := range digitNames {
            if(strings.HasPrefix(str, key)){
                digits = append(digits, key)
            }else if (strings.HasPrefix(str, fmt.Sprint(value))){
                digits = append(digits, fmt.Sprint(value))
            }
        }
    }

    var digits_in_numbers []string 
    for _, val := range digits {
        value, found := digitNames[val]
        if(found){
            digits_in_numbers = append(digits_in_numbers, fmt.Sprint(value))
        }else{
            digits_in_numbers = append(digits_in_numbers, val)
        }
       
    }

    the_digit := createTheDigit(digits_in_numbers)

    return the_digit
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
	
	sum, sum2 := 0, 0;

	for _, line := range lines {
		digit := firstAndLastDigit(line)
        digit2 := firstAndLastDigitString(line)
        sum += digit
        sum2 += digit2
    }

	fmt.Println(sum)
    fmt.Println(sum2)
}
