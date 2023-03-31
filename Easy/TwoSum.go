package main

import "fmt"

/*
Входные данные: числа = [2,7,11,15], цель = 9
Вывод: [0,1]
Объяснение: поскольку числа [0] + числа [1] == 9, мы возвращаем [0, 1]
*/

func main() {
	x := []int{5, 4, 6, 4, 7, 90}
	y := 97
	fmt.Println(twoSum(x, y))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+nums[i+1] == target {
			f := []int{i, i + 1}
			return f
		}
	}
	c := []int{}
	return c
}
