package main

type Solution struct {
}

func Constructor() Solution {
	sol := Solution{}
	return sol
}
func (this *Solution) max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}
func (this *Solution) lengthOfLongestSubString(str string) int32 {
	maxLen, curPos := int32(0), int32(0)
	charPos := make(map[byte]int32, 256)
	var i int32
	for i = 0; i < 256; i++ {
		charPos[byte(i)] = -1
	}
	strLen := int32(len(str))
	for i = 0; i < strLen; i++ {
		prePos := charPos[str[i]]
		if prePos >= curPos {
			curPos = charPos[str[i]]
		}
		charPos[str[i]] = i
		maxLen = this.max(maxLen, i-curPos)
	}
	return maxLen
}

func main() {
	sol := Constructor()
	s := "aabcdd"
	len := sol.lengthOfLongestSubString(s)
	println(len)
}
