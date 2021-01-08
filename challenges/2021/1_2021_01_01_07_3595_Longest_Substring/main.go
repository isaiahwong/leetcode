package main

import "math"

// Sliding window technique
func lengthOfLongestSubstring(s string) int {
	max := 0
	left := 0
	right := 0
	track := make(map[string]bool)
	for right < len(s) {
		key_right := string(s[right])
		key_left := string(s[left])
		if _, ok := track[key_right]; !ok {
			track[key_right] = true
			right++
			max = int(math.Max(float64(max), float64(len(track))))
		} else {
			delete(track, key_left)
			left++
		}
	}
	return max
}
