package main

import "strconv"

type BigNumber struct {
	value string
}

func (n BigNumber) digits() int {
	return len(n.value)
}

func (a BigNumber) add(b BigNumber) (result BigNumber) {
	maxLength := max(len(a.value), len(b.value))
	carry := 0
	result.value = ""

	for i := 0; i < maxLength; i++ {
		aDigit, bDigit := 0, 0
		if i < len(a.value) {
			aDigit = int(a.value[len(a.value)-1-i] - '0')
		}
		if i < len(b.value) {
			bDigit = int(b.value[len(b.value)-1-i] - '0')
		}
		sum := aDigit + bDigit + carry
		carry = sum / 10
		result.value = strconv.Itoa(sum%10) + result.value
	}
	if carry > 0 {
		result.value = "1" + result.value
	}

	return
}
