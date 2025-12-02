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

type item struct {
	rot_dir  string // L or R
	distance int
}

func process_input_file(path string, verbose bool) (int, int, int) {
	dat, err := os.ReadFile(path)
	check(err)

	dat_str := string(dat)
	dat_str_array := strings.Split(dat_str, "\n")
	if verbose {
		fmt.Println("Number of lines read:", len(dat_str_array))
	}

	var num_items int = len(dat_str_array)
	var inputs []item
	for i := range num_items {
		rot_dir := dat_str_array[i][:1]
		distance, err := strconv.Atoi(dat_str_array[i][1:])
		check(err)
		inputs = append(inputs, item{rot_dir: rot_dir, distance: distance})
	}

	if verbose {
		fmt.Println("Number of inputs:", len(inputs))
	}

	var cur_pos int = 50
	var num_hit_zero int = 0
	var num_crosses_zero int = 0
	for _, input := range inputs {
		var mult int
		if input.rot_dir == "L" {
			mult = -1
		} else {
			mult = 1
		}
		var hundreds int = input.distance / 100
		num_crosses_zero += hundreds
		var rest_dist int = input.distance % 100
		var new_pos int = cur_pos + mult*rest_dist
		new_pos = (new_pos%100 + 100) % 100
		if cur_pos != 0 && mult*new_pos < mult*cur_pos {
			num_crosses_zero += 1
		} else if new_pos == 0 {
			num_crosses_zero += 1
		}
		cur_pos = new_pos
		if cur_pos == 0 {
			num_hit_zero += 1
		}
		if cur_pos > 99 || cur_pos < 0 {
			panic(cur_pos)
		}
	}

	return cur_pos, num_hit_zero, num_crosses_zero
}

func main() {

	cur_pos_test, num_hit_zero_test, num_crosses_zero_test := process_input_file("test.txt", false)
	if cur_pos_test != 98 {
		panic(cur_pos_test)
	}
	if num_hit_zero_test != 1 {
		panic(num_hit_zero_test)
	}
	if num_crosses_zero_test != 2 {
		panic(num_crosses_zero_test)
	}
	fmt.Println("Test Passed!")

	cur_pos_test2, num_hit_zero_test2, num_crosses_zero_test2 := process_input_file("test2.txt", false)
	if cur_pos_test2 != 99 {
		panic(cur_pos_test2)
	}
	if num_hit_zero_test2 != 1 {
		panic(num_hit_zero_test2)
	}
	if num_crosses_zero_test2 != 4 {
		panic(num_crosses_zero_test2)
	}
	fmt.Println("Test2 Passed!")

	cur_pos_test_off, num_hit_zero_test_off, num_crosses_zero_test_off := process_input_file("test_off.txt", false)
	if cur_pos_test_off != 32 {
		panic(cur_pos_test_off)
	}
	if num_hit_zero_test_off != 3 {
		panic(num_hit_zero_test_off)
	}
	if num_crosses_zero_test_off != 6 {
		panic(num_crosses_zero_test_off)
	}
	fmt.Println("Official Test Passed!")

	cur_pos, num_hit_zero, num_crosses_zero := process_input_file("input.txt", true)
	fmt.Println("Final position:", cur_pos)
	fmt.Println("Number of times hit zero:", num_hit_zero)
	fmt.Println("Number of times crossed zero:", num_crosses_zero)

}
