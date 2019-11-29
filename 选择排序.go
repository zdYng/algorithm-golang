package main

import "fmt"

func selectSort(slic []int)[]int{
	lens := len(slic)

	for i:=0;i<lens-1;i++{
		// 控制循环轮数
		k:=i
		for j:=i+1;j<lens;j++{
			// 控制比较次数
			if slic[k]>slic[j]{
				// 找出除了slic[k],后面的所有元素中最小的一个元素位置标记
				k=j
			}
		}
		if k!=i{
			// 因为替换了位置标记导致k与i不一样，将标记位置与slic[i]交换
			slic[k],slic[i] = slic[i],slic[k]
		}
	}

	return slic

}

func main(){
	slic := []int{10,8,7,5,6,2,11,16,14,13,1,3}
	fmt.Println(selectSort(slic))
	// [1 2 3 5 6 7 8 10 11 13 14 16]

}
