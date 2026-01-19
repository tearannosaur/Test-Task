package models

import "sort"

type Numbers struct {
	Number int `db:"number" json:"number"`
}

func ValidateNumbers(nums []int) []int {
	sort.Ints(nums)
	return nums
}
