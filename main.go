package main

import (
	"fmt"
	"os"
)

func main() {

	if argsCheck(os.Args) == false {
		fmt.Println("Error")
		return
	}

	var length, width = atoi(os.Args[1]), atoi(os.Args[2])

	fmt.Printf("Length: %d\nWidth: %d\n", length, width)

	// for loops

}

func argsCheck(args []string) bool {

	if len(args) != 3 {
		return false
	}

	if atoi(args[1]) < 1 {
		return false
	}

	if atoi(args[2]) < 1 {
		return false
	}

	return true
}

func atoi(s string) int {

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
