package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init_square_grid(size int) [][]int {
	// initialize 2d distances array efficiently
	data := make([]int, size*size)
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = data[i*size : (i+1)*size]
	}
	return grid
}

func calc_distance_squared(point1 [3]int, point2 [3]int) int {
	return (point1[0]-point2[0])*(point1[0]-point2[0]) + (point1[1]-point2[1])*(point1[1]-point2[1]) + (point1[2]-point2[2])*(point1[2]-point2[2])
}

func connect(box_to_circuit []*[]int, i int, j int) {
	*box_to_circuit[i] = append(*box_to_circuit[i], *box_to_circuit[j]...)
	for _, box := range *box_to_circuit[j] {
		box_to_circuit[box] = box_to_circuit[i]
	}
}

func solve(path string, MAX_NUM_CONNECTIONS int, verbose bool) int {

	dat, err := os.ReadFile(path)
	check(err)

	dat_str := string(dat)
	dat_str_array := strings.Split(dat_str, "\r\n")
	if verbose {
		fmt.Println("Number of lines read:", len(dat_str_array))
	}

	var boxes [][3]int
	for _, line := range dat_str_array {
		inputs_line_str := strings.Split(line, ",")
		var inputs_line_int [3]int
		for i, str := range inputs_line_str {
			value, err := strconv.Atoi(str)
			check(err)
			inputs_line_int[i] = value
		}
		boxes = append(boxes, inputs_line_int)
	}

	distances := init_square_grid(len(boxes))
	// calc distances to fill grid
	for i, box1 := range boxes {
		for j, box2 := range boxes {
			distances[i][j] = calc_distance_squared(box1, box2)
		}
	}

	//for _, ds := range distances {
	//	fmt.Println(ds)
	//}

	box_to_circuit := make([]*[]int, len(boxes)) // circuit is a pointer to a slice of indices of boxes
	// init each circuit with just the box
	for i := range len(boxes) {
		box_to_circuit[i] = &[]int{i}
	}

	// do MAX_NUM_CONNECTIONS
	num_connections := 0
	for num_connections < MAX_NUM_CONNECTIONS {
		// find cur smallest
		cur_smallest := -1
		cur_smallest_idx := [2]int{-1, -1}
		found_unconnected_box := false
		for i := range len(boxes) {
			for j := range len(boxes) {
				if distances[i][j] == 0 {
					continue
				}
				if cur_smallest == -1 || (distances[i][j] < cur_smallest) {
					cur_smallest = distances[i][j]
					cur_smallest_idx = [2]int{i, j}
					found_unconnected_box = true
					continue
				}
			}
		}
		if !found_unconnected_box {
			fmt.Println("no unconnected box found")
			break
		}
		i, j := cur_smallest_idx[0], cur_smallest_idx[1]
		// skip next time
		distances[i][j] = 0
		distances[j][i] = 0
		// check if already in same circuit
		if !slices.Contains(*box_to_circuit[i], j) {
			//fmt.Println("combining", i, j, boxes[i], boxes[j])
			connect(box_to_circuit, i, j)
		} else {
			//fmt.Println("skipping combining", i, j, boxes[i], boxes[j])
		}
		num_connections += 1
	}

	//for k, v := range box_to_circuit {
	//	fmt.Println(k, *v)
	//}

	// get top N biggest circuits
	TOP_N := 3
	top_n := make([]*[]int, TOP_N)
	top_n_sizes := make([]int, TOP_N)
	for i := range len(boxes) {
		cur_circuit_addr := box_to_circuit[i]
		if slices.Contains(top_n, cur_circuit_addr) {
			continue
		}
		size := len(*cur_circuit_addr)
		if size <= top_n_sizes[TOP_N-1] {
			continue
		}
		replace_idx := TOP_N - 1
		for j := TOP_N - 2; j >= 0; j-- {
			if size > top_n_sizes[j] {
				replace_idx = j
				// move down
				top_n_sizes[j+1] = top_n_sizes[j]
				top_n[j+1] = top_n[j]
			} else {
				break
			}
		}
		top_n_sizes[replace_idx] = size
		top_n[replace_idx] = cur_circuit_addr
	}

	//for _, circuit := range top_n {
	//	fmt.Println(circuit)
	//}
	//fmt.Println(top_n_sizes)

	// multiply top n together
	total := top_n_sizes[0]
	for i := 1; i < TOP_N; i++ {
		total = total * top_n_sizes[i]
	}

	return total

}

func solve2(path string, verbose bool) int {

	dat, err := os.ReadFile(path)
	check(err)

	dat_str := string(dat)
	dat_str_array := strings.Split(dat_str, "\r\n")
	if verbose {
		fmt.Println("Number of lines read:", len(dat_str_array))
	}

	var boxes [][3]int
	for _, line := range dat_str_array {
		inputs_line_str := strings.Split(line, ",")
		var inputs_line_int [3]int
		for i, str := range inputs_line_str {
			value, err := strconv.Atoi(str)
			check(err)
			inputs_line_int[i] = value
		}
		boxes = append(boxes, inputs_line_int)
	}

	distances := init_square_grid(len(boxes))
	// calc distances to fill grid
	for i, box1 := range boxes {
		for j, box2 := range boxes {
			distances[i][j] = calc_distance_squared(box1, box2)
		}
	}

	//for _, ds := range distances {
	//	fmt.Println(ds)
	//}

	box_to_circuit := make([]*[]int, len(boxes)) // circuit is a pointer to a slice of indices of boxes
	// init each circuit with just the box
	for i := range len(boxes) {
		box_to_circuit[i] = &[]int{i}
	}

	// do until all together in one circuit
	var last_connection_idxs [2]int
	for len(*box_to_circuit[0]) < len(boxes) {
		// find cur smallest
		cur_smallest := -1
		cur_smallest_idx := [2]int{-1, -1}
		found_unconnected_box := false
		for i := range len(boxes) {
			for j := range len(boxes) {
				if distances[i][j] == 0 {
					continue
				}
				if cur_smallest == -1 || (distances[i][j] < cur_smallest) {
					cur_smallest = distances[i][j]
					cur_smallest_idx = [2]int{i, j}
					found_unconnected_box = true
					continue
				}
			}
		}
		if !found_unconnected_box {
			fmt.Println("no unconnected box found")
			break
		}
		i, j := cur_smallest_idx[0], cur_smallest_idx[1]
		// skip next time
		distances[i][j] = 0
		distances[j][i] = 0
		// check if already in same circuit
		if !slices.Contains(*box_to_circuit[i], j) {
			//fmt.Println("combining", i, j, boxes[i], boxes[j])
			connect(box_to_circuit, i, j)
			last_connection_idxs = [2]int{i, j}
		} else {
			//fmt.Println("skipping combining", i, j, boxes[i], boxes[j])
		}
	}

	return boxes[last_connection_idxs[0]][0] * boxes[last_connection_idxs[1]][0]
}

func main() {

	total_test_off := solve2("test_off.txt", false)
	if total_test_off != 25272 {
		panic(total_test_off)
	}
	fmt.Println("Official Test Passed!")

	total := solve2("input.txt", true)
	fmt.Println("Total:", total)

}
