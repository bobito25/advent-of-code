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

func pop[T any](s *[]T) T {
	// pops last element of slice and returns it
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func remove[T comparable](s []T, x T) []T {
	// remove element from slice and return the slice
	for i := 0; i < len(s); i++ {
		if s[i] == x {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}

func solve(path string, verbose bool) int {
	dat, err := os.ReadFile(path)
	check(err)

	dat_str := string(dat)
	dat_str_array := strings.Split(dat_str, "\r\n\r\n")
	if len(dat_str_array) != 2 {
		panic(len(dat_str_array))
	}
	dat_str_array_ranges := strings.Split(dat_str_array[0], "\r\n")
	dat_str_array_ids := strings.Split(dat_str_array[1], "\r\n")
	if verbose {
		fmt.Println("Number of lines read:", len(dat_str_array_ranges)+len(dat_str_array_ids))
	}

	var input_ranges [][2]int
	max_range_id := 0
	for _, line := range dat_str_array_ranges {
		var range_str []string = strings.Split(line, "-")
		if len(range_str) != 2 {
			panic(range_str)
		}
		range_begin, err := strconv.Atoi(range_str[0])
		check(err)
		range_end, err := strconv.Atoi(range_str[1])
		check(err)
		if range_end < range_begin {
			panic(range_str)
		}
		if range_end > max_range_id {
			max_range_id = range_end
		}
		input_ranges = append(input_ranges, [2]int{range_begin, range_end})
	}
	if verbose {
		fmt.Println("Num ranges:", len(input_ranges))
		fmt.Println("Max range id:", max_range_id)
	}

	var input_ids []int
	for _, id_str := range dat_str_array_ids {
		id, err := strconv.Atoi(id_str)
		check(err)
		input_ids = append(input_ids, id)
	}

	num_fresh := 0
	for _, id := range input_ids {
		for _, id_range := range input_ranges {
			if id >= id_range[0] && id <= id_range[1] {
				num_fresh += 1
				break
			}
		}
	}

	return num_fresh

}

func solve2(path string, verbose bool) int {
	dat, err := os.ReadFile(path)
	check(err)

	dat_str := string(dat)
	dat_str_array := strings.Split(dat_str, "\r\n\r\n")
	if len(dat_str_array) != 2 {
		panic(len(dat_str_array))
	}
	dat_str_array_ranges := strings.Split(dat_str_array[0], "\r\n")
	dat_str_array_ids := strings.Split(dat_str_array[1], "\r\n")
	if verbose {
		fmt.Println("Number of lines read:", len(dat_str_array_ranges)+len(dat_str_array_ids))
	}

	var input_ranges [][2]int
	max_range_id := 0
	for _, line := range dat_str_array_ranges {
		var range_str []string = strings.Split(line, "-")
		if len(range_str) != 2 {
			panic(range_str)
		}
		range_begin, err := strconv.Atoi(range_str[0])
		check(err)
		range_end, err := strconv.Atoi(range_str[1])
		check(err)
		if range_end < range_begin {
			panic(range_str)
		}
		if range_end > max_range_id {
			max_range_id = range_end
		}
		input_ranges = append(input_ranges, [2]int{range_begin, range_end})
	}
	if verbose {
		fmt.Println("Num ranges:", len(input_ranges))
		fmt.Println("Max range id:", max_range_id)
	}

	var combined_ranges [][2]int
	max_iter := 9999
	i := 0
	for len(input_ranges) > 0 {
		i++
		if i > max_iter {
			panic("max iter reached")
		}
		// treat input ranges as stack
		id_range := pop(&input_ranges)
		might_add := true
		var to_remove [][2]int
		for _, com_id_range := range combined_ranges {
			if id_range[0] >= com_id_range[0] && id_range[0] <= com_id_range[1] {
				// beg is in existing range
				if id_range[1] > com_id_range[1] {
					// combine and add to stack
					new_range := [2]int{com_id_range[0], id_range[1]}
					combined_ranges = remove(combined_ranges, com_id_range)
					input_ranges = append(input_ranges, new_range)
					might_add = false
					break
				} else {
					// already contained in existing combined range
					might_add = false
					break
				}
			}
			if id_range[1] >= com_id_range[0] && id_range[1] <= com_id_range[1] {
				// end is in existing range
				if id_range[0] < com_id_range[0] {
					// combine and add to stack
					new_range := [2]int{id_range[0], com_id_range[1]}
					combined_ranges = remove(combined_ranges, com_id_range)
					input_ranges = append(input_ranges, new_range)
					might_add = false
					break
				} else {
					// already contained in existing combined range
					might_add = false
					break
				}
			}
			if id_range[0] < com_id_range[0] && id_range[1] > com_id_range[1] {
				// existing range is inside new range
				to_remove = append(to_remove, com_id_range)
			}
		}
		for _, to_rem_id_range := range to_remove {
			combined_ranges = remove(combined_ranges, to_rem_id_range)
		}
		if might_add {
			combined_ranges = append(combined_ranges, id_range)
		}
	}
	if verbose {
		fmt.Println("Num ranges after combining:", len(combined_ranges))
	}

	num_fresh := 0
	for _, id_range := range combined_ranges {
		num_fresh += id_range[1] - id_range[0] + 1
	}

	return num_fresh

}

func main() {

	num_fresh_test_off := solve2("test_off.txt", false)
	if num_fresh_test_off != 14 {
		panic(num_fresh_test_off)
	}
	fmt.Println("Official Test Passed!")

	num_fresh_test_1 := solve2("test.txt", false)
	if num_fresh_test_1 != 15 {
		panic(num_fresh_test_1)
	}
	fmt.Println("Test 1 Passed!")

	num_fresh_test_2 := solve2("test2.txt", false)
	if num_fresh_test_2 != 15 {
		panic(num_fresh_test_2)
	}
	fmt.Println("Test 2 Passed!")

	num_fresh_test_3 := solve2("test3.txt", false)
	if num_fresh_test_3 != 5 {
		panic(num_fresh_test_3)
	}
	fmt.Println("Test 3 Passed!")

	num_fresh := solve2("input.txt", true)
	fmt.Println("Number of fresh ingredients:", num_fresh)

}
