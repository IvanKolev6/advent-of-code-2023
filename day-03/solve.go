package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func containsSpecialCharacters(s string) bool {
	pattern := `[!@#$%^&*+=~/-]`
	regExp := regexp.MustCompile(pattern)
	matchIndex := regExp.FindStringIndex(s)
	return matchIndex != nil
}

func containsSpecialCharactersContent(grid [][]rune, r_id int, ditigs []int)bool{
	for _,value := range ditigs{
		res := containsSpecialCharacters(string(grid[r_id][value]))
		if res{
			return true
		}
	}
	if ditigs[0] >= 1{
		res := containsSpecialCharacters(string(grid[r_id][ditigs[0]-1]))
		if res{
			return true
		}
	}
	if ditigs[len(ditigs)-1] < len(grid[0])-1{
		res := containsSpecialCharacters(string(grid[r_id][ditigs[len(ditigs)-1]+1]))
		if res{
			return true
		}
	}
	return false
}

func isAdjacent(grid [][]rune, r_id int, ditigs []int)bool{
	if r_id >= 1{
		res := containsSpecialCharactersContent(grid, r_id-1, ditigs)
		if res{
			return true
		}
	}
	if r_id < len(grid)-1{
		res := containsSpecialCharactersContent(grid, r_id+1, ditigs)
		if res{
			return true
		}
	}
	res := containsSpecialCharactersContent(grid, r_id, ditigs)
	if res{
		return true
	}
	return false
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
	
	sum := 0;
	content := make([][]rune, 0)


	for _, line := range lines {
		content = append(content, []rune(line))
    }

	for r_id, row := range content {
		var digit_position  [][]int
		var digit_one  []int
		var digit_2d_slice [][]int	
		var digit_value []int
		row_id := 0
		for pos, char := range row{
			if char_int, err := strconv.Atoi(string(char)); err == nil {
				digit_one = append(digit_one, pos)
				digit_value = append(digit_value, char_int)
			
			}else{
				if len(digit_one) != 0{
					digit_position = append(digit_position, []int{})
					digit_position[row_id] = append(digit_position[row_id], digit_one...)
					digit_one = nil
					digit_2d_slice = append(digit_2d_slice, []int{})
					digit_2d_slice[row_id] = append(digit_2d_slice[row_id], digit_value...)
					digit_value = nil
					row_id++
				}
			}			
		}
		if len(digit_one) != 0{
			digit_position = append(digit_position, []int{})
			digit_position[row_id] = append(digit_position[row_id], digit_one...)
			digit_one = nil
			digit_2d_slice = append(digit_2d_slice, []int{})
			digit_2d_slice[row_id] = append(digit_2d_slice[row_id], digit_value...)
			digit_value = nil
			row_id++
		}

		if len(digit_position) != 0{
			for pos, d_row := range digit_position{
				res := isAdjacent(content, r_id, d_row)
				if res{
					fmt.Println(digit_2d_slice[pos])
					var result string
					for _, num := range digit_2d_slice[pos] {
						result += strconv.Itoa(num)
					}
					res, _ := strconv.Atoi(result);
					sum +=  res
				}
			}
			
		}		
    }
	fmt.Println(sum)
	
}
