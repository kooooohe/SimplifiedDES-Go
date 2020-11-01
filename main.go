package main

import "fmt"

var k1 [8]int
var k2 [8]int

func main() {
	key := [10]int{1, 0, 0, 1, 0, 1, 0, 1, 0, 1}

	rKey := p10(key)
	fmt.Println(rKey)
	rKey = leftShift(rKey)
	fmt.Println(rKey)
	p8(rKey, 1)
	fmt.Println(k1)

	rKey = leftShift(rKey)
	fmt.Println(rKey)
	rKey = leftShift(rKey)
	fmt.Println(rKey)

	p8(rKey, 2)
	fmt.Println(k2)
}

func p10(oKey [10]int) (r [10]int) {
	p10 := [10]int{2, 4, 1, 6, 3, 9, 0, 8, 7, 5}
	for k, v := range p10 {
		r[k] = oKey[v]
	}
	return
}

func leftShift(key [10]int) [10]int {
	tmp := key[0]
	for i := 0; i < 4; i++ {
		key[i] = key[i+1]
	}
	key[4] = tmp
	tmp = key[5]

	for i := 5; i < 9; i++ {
		key[i] = key[i+1]
	}
	key[9] = tmp
	return key
}

func p8(key [10]int, k int) {
	p8 := [8]int{5, 2, 6, 3, 7, 4, 9, 8}
	if k == 1 {
		for k, v := range p8 {
			k1[k] = key[v]
		}
	} else {
		for k, v := range p8 {
			k2[k] = key[v]
		}
	}
}
