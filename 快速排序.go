package main

import "fmt"

func quickSort(slic []int) []int{
	lens := len(slic)

	if lens<=1{
		return slic
	}

	select_num := slic[0]
	left_slic := []int{}
	right_slic := []int{}
	for i := 1;i<lens;i++{
		if select_num>slic[i]{
			left_slic = append(left_slic,slic[i])
		}else{
			right_slic = append(right_slic, slic[i])
		}
	}

	left_slic = quickSort(left_slic)
	right_slic = quickSort(right_slic)



	left_slic = append(left_slic, select_num)

	fmt.Println("left:",left_slic)
	fmt.Println("right:",right_slic)

	return append(left_slic, right_slic...)

}

func main(){
	slic := []int{6,4,3,1,2,5,10,9,8,7}
	res := quickSort(slic)
	fmt.Println(res)
}


