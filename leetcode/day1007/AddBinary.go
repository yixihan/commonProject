package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(addBinary("1111", "10101"))
}

/*
    1 1 1 1
  1 0 1 0 1
1 0 0 1 0 0
*/

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	sb := new(strings.Builder)
	tmp := 0
	for i >= 0 || j >= 0 {
		val := 0
		if i >= 0 {
			val += int(a[i]) - 48
			i--
		}
		if j >= 0 {
			val += int(b[j]) - 48
			j--
		}
		if tmp > 0 {
			val += tmp
			tmp = 0
		}
		if val > 1 {
			tmp = 1
			val -= 2
		}
		sb.WriteString(strconv.Itoa(val))
		fmt.Println(val, sb)
	}
	sb.WriteString(strconv.Itoa(tmp))
	fmt.Printf("sb = %v\n", sb)
	return reverse(sb.String())
}

func reverse(sb string) string {
	ans := new(strings.Builder)
	flag := false
	for i := len(sb) - 1; i >= 0; i-- {
		if !flag && sb[i] == '0' {
			continue
		} else if sb[i] != '0' {
			flag = true
			ans.WriteString(string(sb[i]))
		} else {
			ans.WriteString(string(sb[i]))
		}
	}
	if len(ans.String()) == 0 {
		ans.WriteString("0")
	}
	return ans.String()
}
