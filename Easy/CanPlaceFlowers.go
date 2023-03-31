package main

import (
	"fmt"
)

/*
У вас длинная клумба, на которой некоторые участки засажены, а некоторые нет.
Однако цветы нельзя сажать на соседних участках.

Учитывая целочисленный массивflowerbed, содержащий 0's и 1's,
где 0означает пустой и 1означает не пустой, и целое числоn, возвращает,
могут ли n новые цветы быть посажены вflowerbed, не нарушая правило отсутствия смежных цветов.

Ввод: клумба = [1,0,0,0,1], n = 2
Вывод: false
*/

func main() {
	var a []int = []int{1, 0, 0, 0, 0, 1}
	n := 3
	fmt.Println(canPlace(a, n))
}

func canPlace(clumb []int, n int) bool {
	if len(clumb) > 1000 {
		return false
	}
	if n > len(clumb) {
		return false
	}
	for i := 0; i <= len(clumb)-1; i++ {
		if clumb[i] > 1 || clumb[i] < 0 {
			return false
		} else {
			return true
		}
		if clumb[i] == 0 && clumb[i+1] == 0 && clumb[i+2] == 0 && clumb[i-1] != 1 {
			return false
		}
	}
	return false
}
