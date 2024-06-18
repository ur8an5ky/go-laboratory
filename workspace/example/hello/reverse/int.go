package reverse

import "strconv"

func Int(num int) int {
	num, _ = strconv.Atoi(String(strconv.Itoa(num)))
	return num
}
