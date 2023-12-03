package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getTheGameID(line string) (int, int){
	gameReq := map[string]int{
		"red" : 12,
		"green" : 13,
		"blue" : 14,
	}
	p1 := true
	p2 := make(map[string]int)
	games := strings.Split(line, ":")
	game := strings.Split(games[1], ";")
	for _, val := range game{
		cubes := strings.Split(val, ",")
		var cube_id []string
		var cube_color []string
		for _, val1 := range cubes{
			cubes_eq := strings.Split(val1, " ")
			cube_id = append(cube_id, cubes_eq[1])
			cube_color = append(cube_color, cubes_eq[2])
		}

		for pos := range cube_color{
			value := gameReq[cube_color[pos]]
			current_id, _ :=  strconv.Atoi(cube_id[pos])
			p2[cube_color[pos]] = max(p2[cube_color[pos]], current_id)
			if value < current_id{
				p1 = false
			}			
		}

	}

	score := 1
	for _, val := range p2 {
		score *= val
	}
	game_id_int, _ :=  strconv.Atoi(strings.Split(games[0], " ")[1])
	if p1 == false{
		return 0, score
	}
	return game_id_int, score
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
		id, score := getTheGameID(line)
        sum += id
		sum2 += score
    }

	fmt.Println(sum)
	fmt.Println(sum2)
}
