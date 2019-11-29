package main

import "fmt"

func bubbleSort(slic []int)[]int{
	lens := len(slic)

	for i:=1;i<lens;i++{
		for j:=0;j<lens-1;j++{
			if slic[i]<slic[j]{
				slic[i],slic[j]=slic[j],slic[i] // 前面元素大于后面元素则交换位置
			}
		}
	}
	return slic
}

func main(){
	slic := []int{3,1,2,6,5,4,8,7,9}
	res := bubbleSort(slic)
	fmt.Println(res)
}

//运行结果
// [1 2 3 4 5 6 7 8 9]