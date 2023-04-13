package main

import (
	"fmt"
	//"strconv"
)

func main() {
	var str string
	str = "him"
	var t =[]byte(str)
	t[0] = str[2]
	fmt.Printf("%v, %T", string(t), t)

}
