package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func solve(path string, verbose bool) int {

	dat, err := os.ReadFile(path)
	check(err)

	dat_str := string(dat)
	dat_str_array := strings.Split(dat_str, "\r\n")
	if verbose {
		fmt.Println("Number of lines read:", len(dat_str_array))
	}

	var inputs_str [][]string
	for _, line := range dat_str_array {
		inputs_str = append(inputs_str, strings.Split(line, ""))
	}

	start := strings.Index(dat_str_array[0], "S")
	if start == -1 {
		panic("No S found")
	}

	// init
	var cur_beams []int
	for i := range len(inputs_str[0]) {
		if i == start {
			cur_beams = append(cur_beams, 1)
		} else {
			cur_beams = append(cur_beams, 0)
		}
	}

	// iterate through rows
	num_split := 0
	for i := 1; i < len(dat_str_array); i++ {
		new_beams := cur_beams
		for j, cell_underneath := range inputs_str[i] {
			if cur_beams[j] == 1 {
				if cell_underneath == "^" {
					// split
					new_beams[j-1], new_beams[j+1] = 1, 1
					new_beams[j] = 0
					num_split += 1
				} else {
					new_beams[j] = 1
				}
			}
		}
		cur_beams = new_beams
	}

	return num_split

}

func solve2(path string, verbose bool) int {

	dat, err := os.ReadFile(path)
	check(err)

	dat_str := string(dat)
	dat_str_array := strings.Split(dat_str, "\r\n")
	if verbose {
		fmt.Println("Number of lines read:", len(dat_str_array))
	}

	var inputs_str [][]string
	for _, line := range dat_str_array {
		inputs_str = append(inputs_str, strings.Split(line, ""))
	}

	start := strings.Index(dat_str_array[0], "S")
	if start == -1 {
		panic("No S found")
	}

	// init
	var cur_beams []int
	for i := range len(inputs_str[0]) {
		if i == start {
			cur_beams = append(cur_beams, 1)
		} else {
			cur_beams = append(cur_beams, 0)
		}
	}

	// iterate through rows
	for i := 1; i < len(dat_str_array); i++ {
		new_beams := cur_beams
		for j, cell_underneath := range inputs_str[i] {
			if cur_beams[j] >= 1 {
				if cell_underneath == "^" {
					// split
					new_beams[j-1] += new_beams[j]
					new_beams[j+1] += new_beams[j]
					new_beams[j] = 0
				} else {
					new_beams[j] = cur_beams[j]
				}
			}
		}
		cur_beams = new_beams
	}

	num_timelines := 0
	for _, num_end_here := range cur_beams {
		num_timelines += num_end_here
	}

	return num_timelines

}

func main() {

	num_timelines_test_off := solve2("test_off.txt", false)
	if num_timelines_test_off != 40 {
		panic(num_timelines_test_off)
	}
	fmt.Println("Official Test Passed!")

	num_timelines := solve2("input.txt", true)
	fmt.Println("Total timelines:", num_timelines)

}
