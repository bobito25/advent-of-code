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

func solve(path string, verbose bool) (int, int) {
	dat, err := os.ReadFile(path)
	check(err)

	dat_str := string(dat)
	dat_str_array := strings.Split(dat_str, ",")
	if verbose {
		fmt.Println("Number of items read:", len(dat_str_array))
	}

	var inputs [][2]int

	for _, item := range dat_str_array {
		range_str := strings.Split(item, "-")
		if len(range_str) != 2 {
			panic(range_str)
		}
		beg, err := strconv.Atoi(range_str[0])
		check(err)
		end, err := strconv.Atoi(range_str[1])
		check(err)
		range_arr := [2]int{beg, end}
		inputs = append(inputs, range_arr)
	}

	var invalid_ids []int

	for _, id_range := range inputs {
		beg, end := id_range[0], id_range[1]
		for i := beg; i <= end; i++ {
			str_i := strconv.Itoa(i)
			i_len := len(str_i)
			if i_len%2 == 0 {
				div := i_len / 2
				if str_i[:div] == str_i[div:] {
					invalid_ids = append(invalid_ids, i)
				}
			}
		}
	}

	num_invalid := len(invalid_ids)
	var invalid_sum int = 0
	for _, id := range invalid_ids {
		invalid_sum += id
	}

	return num_invalid, invalid_sum
}

func solve2(path string, verbose bool) (int, int) {
	dat, err := os.ReadFile(path)
	check(err)

	dat_str := string(dat)
	dat_str_array := strings.Split(dat_str, ",")
	if verbose {
		fmt.Println("Number of items read:", len(dat_str_array))
	}

	var inputs [][2]int

	for _, item := range dat_str_array {
		range_str := strings.Split(item, "-")
		if len(range_str) != 2 {
			panic(range_str)
		}
		beg, err := strconv.Atoi(range_str[0])
		check(err)
		end, err := strconv.Atoi(range_str[1])
		check(err)
		range_arr := [2]int{beg, end}
		inputs = append(inputs, range_arr)
	}

	var invalid_ids []int

	for _, id_range := range inputs {
		beg, end := id_range[0], id_range[1]
		for i := beg; i <= end; i++ {
			str_i := strconv.Itoa(i)
			i_len := len(str_i)
			for pat_len := 1; pat_len <= i_len/2; pat_len++ {
				if str_i[pat_len:]+str_i[:pat_len] == str_i {
					invalid_ids = append(invalid_ids, i)
					fmt.Println(str_i[pat_len:], str_i[:pat_len], str_i)
					break
				}
			}
		}
	}

	num_invalid := len(invalid_ids)
	var invalid_sum int = 0
	for _, id := range invalid_ids {
		invalid_sum += id
	}

	return num_invalid, invalid_sum
}

func main() {
	num_invalid_test_off, invalid_sum_test_off := solve("test_off.txt", false)
	if num_invalid_test_off != 8 {
		panic(num_invalid_test_off)
	}
	if invalid_sum_test_off != 1227775554 {
		panic(invalid_sum_test_off)
	}

	num_invalid_test_off2, invalid_sum_test_off2 := solve2("test_off.txt", false)
	if num_invalid_test_off2 != 13 {
		panic(num_invalid_test_off2)
	}
	if invalid_sum_test_off2 != 4174379265 {
		panic(invalid_sum_test_off2)
	}

	num_invalid, invalid_sum := solve2("input.txt", true)
	fmt.Println("Found", num_invalid, "invalid ids!")
	fmt.Println("Total sum:", invalid_sum)
}
