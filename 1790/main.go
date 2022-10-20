package main

import "fmt"

/**
title: Check if One String Swap Can Make Strings Equal
You are given two strings s1 and s2 of equal length. A string swap is an operation where you choose two indices in a string (not necessarily different) and swap the characters at these indices.
Return true if it is possible to make both strings equal by performing at most one string swap on exactly one of the strings. Otherwise, return false.
*/
func main() {
	fmt.Println(areAlmostEqual("zaa", "aaz"))
}

func areAlmostEqual(s1 string, s2 string) bool {
	s1Len, s2Len, cnt := len(s1), len(s2), 0
	if s1Len != s2Len {
		return false
	}
	if s1 == s2 {
		return true
	}
	var tempS1, tempS2 uint8
	for i := 0; i < s1Len; i++ {
		if s1[i] != s2[i] {
			cnt++
			if cnt == 1 {
				tempS1 = s1[i]
				tempS2 = s2[i]
			} else if cnt == 2 {
				if tempS1 != s2[i] || tempS2 != s1[i] {
					return false
				}
			} else {
				return false
			}
		}
	}

	return cnt == 0 || cnt == 2
}
