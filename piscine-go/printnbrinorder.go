package piscine

import "github.com/01-edu/z01"

// execution part
func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	arr := IntArray(n)
	count := 0
	for range arr {
		count++
	}
	sort := Sorter(arr, count)
	for _, v := range sort {
		z01.PrintRune(rune(v) + '0')
	}
} // functions for main function

func Sorter(nums []int, pl int) []int {
	for i := 0; i < pl; i++ {
		for n := i; n < pl; n++ {
			if nums[i] > nums[n] {
				// like the switch assignment
				nums[i], nums[n] = nums[n], nums[i]
			}
		}
	}
	return nums
}

func IntArray(n int) []int {
	if n == 0 {
		return []int{}
	}
	// val
	val := int(n / 10)
	num := int(n % 10)
	return append(IntArray(val), num)
}
