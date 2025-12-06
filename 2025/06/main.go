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

func split_by_spaces(str string) []string {
	str = strings.Trim(str, " ")
	for strings.Contains(str, "  ") {
		str = strings.ReplaceAll(str, "  ", " ")
	}
	str_array := strings.Split(str, " ")
	return str_array
}

func solve(path string, verbose bool) int {
	dat, err := os.ReadFile(path)
	check(err)

	dat_str := string(dat)
	dat_str_array := strings.Split(dat_str, "\r\n")
	if verbose {
		fmt.Println("Number of lines read:", len(dat_str_array))
	}

	var input_equations_str [][]string
	for range len(split_by_spaces(dat_str_array[0])) {
		input_equations_str = append(input_equations_str, []string{})
	}
	for _, line := range dat_str_array {
		split_line := split_by_spaces(line)
		for i, str := range split_line {
			input_equations_str[i] = append(input_equations_str[i], str)
		}
	}

	total_sum := 0
	for _, equation_array := range input_equations_str {
		cur_total, err := strconv.Atoi(equation_array[0])
		check(err)
		op := equation_array[len(equation_array)-1]
		for i := 1; i < len(equation_array)-1; i++ {
			cur_operand, err := strconv.Atoi(equation_array[i])
			check(err)
			switch op {
			case "*":
				cur_total = cur_total * cur_operand
			case "+":
				cur_total = cur_total + cur_operand
			default:
				panic(op)
			}
		}
		total_sum += cur_total
	}

	return total_sum

}

func solve2(path string, verbose bool) int {
	return 0
}

func main() {

	total_sum_test_off := solve("test_off.txt", false)
	if total_sum_test_off != 4277556 {
		panic(total_sum_test_off)
	}
	fmt.Println("Official Test Passed!")

	total_sum := solve("input.txt", true)
	fmt.Println("Total sum:", total_sum)

}
