package main

import (
	//"fmt"
	"strconv"
	//s "strings"
)

func FakeBin(x string) (result string) {
	runes := []rune(x)
	for i, r := range runes {
		num, _ := strconv.Atoi(string(r))
		if num < 5 {
			runes[i] = '0'
		} else {
			runes[i] = '1'
		}
	}
	result = string(runes)
	//fmt.Println(result)
	return result
}

/*func main() {
	FakeBin("123456789")
}

Given a string of digits, you should replace any digit below 5 with '0' and
any digit 5 and above with '1'. Return the resulting string.

func FakeBin(x string) (result string) {
    for _, char := range x {
        if char < '5' {
            result += "0"
        } else {
            result += "1"
        }
    }
    return
}


func FakeBin(x string) (r string) {

	res := strings.NewReplacer("1", "0", "2", "0", "3", "0", "4", "0", "5", "1", "6", "1", "7", "1", "8", "1", "9", "1")

	r = res.Replace(x)
	return
}

*/
