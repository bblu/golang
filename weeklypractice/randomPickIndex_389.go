package main

import (
	"fmt"
	"math/rand"
)
type Solution struct{
	nums []int
}

func Constructor(nums []int) Solution{
	sol := Solution{nums:nums}
	return sol
}

func (this *Solution) pickIndex(target int) int{
	lng := len(this.nums)
	idx := -1
	cnt := 0
	for i:=0; i<lng; i++{
		if this.nums[i] == target{
			cnt += 1
			if rand.Intn(cnt) == 0{
				idx = i
			}
		}
	}
	return idx
}

func main() {
	nums :=[6]int{1, 2, 3, 3, 3, 3}
	obj := Constructor(nums[:])
	for i:=0; i < 5; i++ {
		param_1 := obj.pickIndex(3)
		fmt.Println("pick(3) index =", param_1)
	}
}