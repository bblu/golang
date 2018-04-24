package main

import "fmt"

func getIndicesByTarget(nums []int, leng int, tar int)(i0 int,i1 int){
	for i:=0; i< leng-1; i++{
		p0, p1 := -1, -1
		for j:=i; j< leng; j++{
			fmt.Println(i,j,nums[j])
			if p0 < 0 && nums[j] < tar{
				p0 = i + j
				p1 = tar - nums[j]
				fmt.Println("find index0 : ", i+j, p0, p1)
			}else if nums[j] == p1{
				p1 = i + j
				fmt.Println("find index1 : ", p0, p1)
				return p0, p1
			}
		}
	}
	return -1, -1
}

func main() {
	nums :=[6]int{1, 2, 10, 7, 5, 8}
	tar := 7 
	a,b := getIndicesByTarget(nums[:], 6, tar)
	fmt.Println("vim-go", a, b)
}

