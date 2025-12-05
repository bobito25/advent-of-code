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

func check_exists(grid [][]int, i int, j int) bool {
	if i < 0 || i >= len(grid) {
		return false
	}
	if j < 0 || j >= len(grid[i]) {
		return false
	}
	return true
}

func solve(path string, verbose bool) int {
	dat, err := os.ReadFile(path)
	check(err)

	dat_str := string(dat)
	dat_str_array := strings.Split(dat_str, "\r\n")
	if verbose {
		fmt.Println("Number of lines read:", len(dat_str_array))
	}

	var inputs [][]int
	for _, line := range dat_str_array {
		var line_str []string = strings.Split(line, "")
		var line_int []int
		for _, str := range line_str {
			if str == "@" {
				line_int = append(line_int, 1)
			} else {
				line_int = append(line_int, 0)
			}
		}
		inputs = append(inputs, line_int)
	}

	num_accessible := 0
	var removed_idxs [][2]int
	peturbations := [8][2]int{{1, 1}, {1, 0}, {1, -1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, 1}, {0, -1}}
	stop := false
	for stop == false {
		var to_remove_idxs [][2]int
		for i, row := range inputs {
			for j, cell := range row {
				if cell == 0 {
					continue
				}
				num_nearby := 0
				for _, peturb := range peturbations {
					i_p, j_p := i+peturb[0], j+peturb[1]
					if check_exists(inputs, i_p, j_p) {
						num_nearby += inputs[i_p][j_p]
					}
				}
				if num_nearby < 4 {
					num_accessible += 1
					to_remove_idxs = append(to_remove_idxs, [2]int{i, j})
					continue
				}
			}
		}
		removed_idxs = append(removed_idxs, to_remove_idxs...)
		for _, idxs := range to_remove_idxs {
			inputs[idxs[0]][idxs[1]] = 0
		}
		if len(to_remove_idxs) == 0 {
			stop = true
		}
	}

	return num_accessible

}

func main() {

	num_accessible_test_off := solve("test_off.txt", false)
	if num_accessible_test_off != 43 {
		panic(num_accessible_test_off)
	}
	fmt.Println("Official Test Passed!")

	num_accessible := solve("input.txt", true)
	fmt.Println("Number of accessible rolls:", num_accessible)

}
