package main

import "fmt"

func main() {
	var numbers [5]int // allocates memory for 5 numbers to store.
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	// fmt.Println(numbers)

	var elems [5]interface{} // arrays using the go interface.
	elems[0] = 1
	elems[1] = "John doe"
	elems[2] = 4.5
	elems[3] = true
	elems[4] = 'A'
	// fmt.Println(elems)

	items := make([]int, 5) // arrays using make function
	items[0] = 5
	items[1] = 10
	items[2] = 20
	items[3] = 30
	items[4] = 40
	// fmt.Println(items)

	nums := [5]int{1, 2, 3, 4, 5} // iterating array
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// Exercises :
	// 1) Given an array of integers, find the **maximum and minimum** values.
	MinMax := func(arr []int) (int, int) {
		if len(arr) == 0 {
			panic("array cannot be empty")
		}
		min := arr[0]
		max := arr[0]

		for _, v := range arr {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		return min, max
	}

	arr := [5]int{100, -1, 5000, 99, 17}
	min, max := MinMax(arr[:])
	fmt.Println("Min : ", min)
	fmt.Println("Max : ", max)

	// 2) Given an array, **reverse it in-place**
	reverse := func(arr []int) {
		n := len(arr)

		for i := 0; i < n/2; i++ {
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
	}
	e := []int{1, 2, 3, 4, 5, 6}
	reverse(e)
	// fmt.Println(e)

	// 3) Write a program that:
	// (Passes an array to a function, Modifies its contents correctly, Demonstrates the difference between value and reference behavior)
	modifyByValue := func(arr [5]int) {
		arr[0] = 100
		fmt.Println("Inside modifyByValue : ", arr)
	}

	modifyByReference := func(arr *[5]int) {
		arr[0] = 200
		fmt.Println("Inside modifyByReference : ", *arr)
	}

	modifySlice := func(slc []int) {
		slc[0] = 300
		fmt.Println("Inside modifySlice : ", slc)
	}

	rev := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Original array:", rev)

	// 1. Pass by value
	modifyByValue(rev)
	fmt.Println("After modifyByValue:", rev)

	fmt.Println("------------")

	// 2. Pass by reference (pointer)
	modifyByReference(&rev)
	fmt.Println("After modifyByReference:", rev)

	fmt.Println("------------")

	// 3. Slice behavior
	slice := rev[:] // slice referencing same array
	modifySlice(slice)
	fmt.Println("After modifySlice:", rev)

}
