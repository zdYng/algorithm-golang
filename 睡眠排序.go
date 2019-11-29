package main

import (
	"fmt"
	"time"
)

func sleepSort(slic []int) []int{

	ch := make(chan int)

	for _,value := range slic{
		go func(val int) {
			time.Sleep(time.Duration(val) * 1000000)
			fmt.Println(val)
			ch <- val
		}(value)
	}
	res := []int{}
	for _ = range slic{
		res = append(res,<- ch)
	}
	return res
}

func main(){
	slic := []int{11,9,8,10,6,5,2,1,3,7,4}
	fmt.Println(sleepSort(slic))

	//1
	//2
	//3
	//4
	//5
	//7
	//6
	//9
	//8
	//10
	//11
	//[1 2 3 4 5 7 6 9 8 10 11]
}

