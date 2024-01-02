package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanMap = []struct {
	decVal int
	symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func decimalToRoman(num int) string {
	if num == 0 {
		return ""
	}
	for _, pair := range romanMap {
		if num >= pair.decVal {
			return pair.symbol + decimalToRoman(num-pair.decVal)
		}
	}
	return ""
}

func implContains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

func Decode(roman string) int {
	var sum int
	var Roman = map[byte]int{'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}
	for k, v := range roman {
		if k < len(roman)-1 && Roman[byte(roman[k+1])] > Roman[byte(roman[k])] {
			sum -= Roman[byte(v)]
		} else {
			sum += Roman[byte(v)]
		}
	}
	return sum
}

func main() {

	var count1 int
	var count2 int
	c1 := 0
	c2 := 0
	fl_Roman := 0
	Roman := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	greek := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	answer := 0

	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')

	object := strings.Fields(expression)

	if len(object) != 3 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		os.Exit(1)
	}

	if object[1] != "+" && object[1] != "-" && object[1] != "/" && object[1] != "*" {
		fmt.Println("Вывод ошибки, так как оператор (+, -, /, *) неверен или не на своём месте")
		os.Exit(1)
	}

	operand := object[1]

	for i := 0; i < 3; i++ {
		if implContains(Roman, object[i]) {
			c1++
		}
		if implContains(greek, object[i]) {
			c2++
		}
	}

	if c1 == 2 {
		count1 = Decode(object[0])
		count2 = Decode(object[2])
		fl_Roman = 1
	}

	if (c1 == 2 && c2 == 2) || (c1 != 2 && c2 != 2) {
		fmt.Println("Вывод ошибки, так как — должно быть два операнда. Оба от 1 до 10 или оба от I до X")
		os.Exit(1)
	}

	if fl_Roman == 0 {
		count3, err1 := strconv.Atoi(object[0])
		count4, err2 := strconv.Atoi(object[2])

		if err1 != nil || err2 != nil {
			fmt.Println("Вывод ошибки, так как вы не коректно ввели число(а)")
			os.Exit(1)
		} else {
			count1 = count3
			count2 = count4
		}
	}

	switch operand {
	case "+":
		answer = count1 + count2
	case "-":
		answer = count1 - count2
	case "*":
		answer = count1 * count2
	case "/":
		answer = count1 / count2
	}

	if c1 == 2 && answer <= 0 {
		fmt.Println("Вывод ошибки, так как в римских числах нет нуля и отрицательных чисел")
		os.Exit(1)
	}

	if c1 == 2 {
		fmt.Println(decimalToRoman(answer))
	} else {
		fmt.Println(answer)
	}
}
