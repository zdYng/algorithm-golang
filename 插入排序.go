package main

import "fmt"

func insertSort(slic []int)[]int{
	lens := len(slic)

	for i := 0; i<lens;i++{
		tmp := slic[i]

		for j := i-1;j>=0;j--{
			if tmp<slic[j]{
				slic[j+1], slic[j] = slic[j],tmp
			}else{
				break
			}
		}
	}
	return slic
}

func main(){
	slic := []int{10,6,4,3,5,9,2,1,8,7}
	fmt.Println(insertSort(slic))
	// [1 2 3 4 5 6 7 8 9 10]
}

