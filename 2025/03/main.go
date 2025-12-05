package main

import (
	"fmt"
	"os"
	"strconv"
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

	var inputs_int [][]int
	for _, strs := range inputs_str {
		var inputs_int_line []int
		for _, str := range strs {
			i_int, err := strconv.Atoi(str)
			check(err)
			inputs_int_line = append(inputs_int_line, i_int)
		}
		inputs_int = append(inputs_int, inputs_int_line)
	}

	var total_max_joltage int = 0
	for i, str_bank := range inputs_str {
		int_bank := inputs_int[i]
		max_j1_idx := 0
		for i := 1; i < len(int_bank)-1; i++ {
			if int_bank[i] > int_bank[max_j1_idx] {
				max_j1_idx = i
			}
		}
		max_j2_idx := max_j1_idx + 1
		for i := max_j1_idx + 1; i < len(int_bank); i++ {
			if int_bank[i] > int_bank[max_j2_idx] {
				max_j2_idx = i
			}
		}
		combined_int, err := strconv.Atoi(str_bank[max_j1_idx] + str_bank[max_j2_idx])
		check(err)
		total_max_joltage += combined_int
	}

	return total_max_joltage
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

	var inputs_int [][]int
	for _, strs := range inputs_str {
		var inputs_int_line []int
		for _, str := range strs {
			i_int, err := strconv.Atoi(str)
			check(err)
			inputs_int_line = append(inputs_int_line, i_int)
		}
		inputs_int = append(inputs_int, inputs_int_line)
	}

	const NUM_ACTIVE_BATTERIES_PER_BANK int = 12

	var total_max_joltage int = 0
	for i, str_bank := range inputs_str {
		int_bank := inputs_int[i]
		var max_j_idxs []int
		for num_bat := range NUM_ACTIVE_BATTERIES_PER_BANK {
			cur_max_idx := 0
			if len(max_j_idxs) > 0 {
				cur_max_idx = max_j_idxs[num_bat-1] + 1
			}
			to_go := NUM_ACTIVE_BATTERIES_PER_BANK - (num_bat + 1)
			until := len(int_bank) - to_go
			for idx := cur_max_idx + 1; idx < until; idx++ {
				if int_bank[idx] > int_bank[cur_max_idx] {
					cur_max_idx = idx
				}
			}
			max_j_idxs = append(max_j_idxs, cur_max_idx)
		}
		combined_str := ""
		for _, idx := range max_j_idxs {
			combined_str += str_bank[idx]
		}
		combined_int, err := strconv.Atoi(combined_str)
		check(err)
		total_max_joltage += combined_int
	}

	return total_max_joltage
}

func main() {

	max_joltage_test_off := solve("test_off.txt", false)
	if max_joltage_test_off != 357 {
		panic(max_joltage_test_off)
	}
	fmt.Println("Official Test Passed!")

	max_joltage_test_off2 := solve2("test_off.txt", false)
	if max_joltage_test_off2 != 3121910778619 {
		panic(max_joltage_test_off2)
	}
	fmt.Println("Official Test Passed!")

	max_joltage := solve2("input.txt", true)
	fmt.Println("Max Joltage:", max_joltage)

}
