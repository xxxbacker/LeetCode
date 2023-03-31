package main

import (
	"fmt"
	"strings"
)

/*
Символ значения
I 1
V 5
X 10
L 50
C 100
D 500
M 1000

Учитывая римскую цифру, преобразуйте ее в целое число.

Ввод: s = "MCMXCIV"
Вывод: 1994
Объяснение: M = 1000, CM = 900, XC = 90 и IV = 4

1 <= s.length <= 15
s содержит только символы ('I', 'V', 'X', 'L', 'C', 'D', 'M').
Гарантируется, что s это допустимая римская цифра в диапазоне [1, 3999]
*/
func main() {
	fmt.Println(romeToInt("IIII"))
}

func romeToInt(s string) int {
	var sum int
	arrS := strings.Split(s, "")
	if len(s) < 1 || len(s) > 15 {
		fmt.Println("err")
		return 228
	}
	/*for i := 0; i < len(arrS)-1; i++ {
	if arrS[i] != "I" || arrS[i] != "V" || arrS[i] != "X" || arrS[i] != "L" || arrS[i] != "C" || arrS[i] != "D" || arrS[i] != "M" {
		fmt.Println("err")
		return 228
	}
	*/
	//}

	rome := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for key, _ := range rome {
		if _, ok := rome[key]; ok {
			for ok == false {
				fmt.Println(228)
			}
		}
	}
	var i = 0
	for i < len(arrS)-1 {
		switch {
		case rome[arrS[i]] == rome[arrS[i+1]] && rome[arrS[i]] == rome[arrS[i+2]]:
			sum += rome[arrS[i]] + rome[arrS[i+1]] + rome[arrS[i+2]]
			i += 3
		case rome[arrS[i]] == rome[arrS[i+1]]:
			sum += rome[arrS[i]] + rome[arrS[i+1]]
			i += 2
		case rome[arrS[i]] > rome[arrS[i+1]]:
			sum += rome[arrS[i]]
			i += 1
		case rome[arrS[i]] < rome[arrS[i+1]]:
			sum += rome[arrS[i+1]] - rome[arrS[i]]
			i += 2
		}
	}
	return sum
}
