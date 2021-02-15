package main

import (
	"fmt"
	"os"
)

func main() {

	if argsValid(os.Args) == false {
		PrintStr("Error")
		return
	}

	var length, width = Atoi(os.Args[1]), Atoi(os.Args[2])

	var white = " "
	var black = "#"
	var newLine = "\n"

	var prev = black

	for i := 1; i <= width; i++ {

		if prev == white {
			prev = black
		} else if prev == black {
			prev = white
		}

		for j := 1; j <= length; j++ {

			if prev == white {
				fmt.Print(black)
				prev = black
			} else if prev == black {
				fmt.Print(white)
				prev = white
			}

		}

		fmt.Print(newLine)
	}

}

func argsValid(args []string) bool {

	if len(args) != 3 {
		return false
	}

	if Atoi(args[1]) < 1 {
		return false
	}

	if Atoi(args[2]) < 1 {
		return false
	}

	return true
}

func Atoi(s string) int {

	var f = false
	var signCount, str = 0, 0

	for i, v := range s {
		// 43 = '+', 45 = '-', 48 = '0', 57 = '9'
		if v == 43 || v == 45 || v >= 48 && v <= 57 {

			if (v == 43 || v == 45) && i > 0 {
				return 0
			}

			if v == 45 && i == 0 {
				f = true
			}

			if v == 43 || v == 45 {
				signCount++
				if signCount > 1 {
					return 0
				}
				continue
			}

			var a = 0
			var j int32
			for j = 49; j <= v; j++ {
				a++
			}
			str = str*10 + a

		} else {
			return 0
		}

	}

	if f == true {
		str = str - (str * 2)
	}

	return str
}

func PrintStr(s string) {
	for _, v:= range s {
		fmt.Print(v)
	}
}
