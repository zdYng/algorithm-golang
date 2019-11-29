# algorithm-golang

### 冒泡排序
一个无序的数组序列，通过从前往后对比相邻的两个元素并调整，让较大的数往后排，较小的数往前排，就像吐泡泡一样。
```golang
package main

import "fmt"

func bubbleSort(slic []int)[]int{
	lens := len(slic)

	for i:=1;i<lens;i++{
		for j:=0;j<lens-1;j++{
			if slic[i]<slic[j]{
				slic[i],slic[j]=slic[j],slic[i] 
				// 前面元素大于后面元素则交换位置
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
```

### 选择排序

应该算是冒泡排序的修改吧，在一个无序的数组中，选出最小的元素与第一个元素位换位置，然后在剩余的数组继续操作，如此循环到倒数第二个元素与最后一个元素比较交换为止。

``` golang
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

```


### 快速排序

在一个无序数组序列中，选择一个元素（一般为第一个或者最后一个元素），通过一轮比较，将该数组序列分成两部分，一部分为小于选择的元素，一部分为大于等于选择的元素。此时该选择元素为正确的位置，然后以递归方式，再重复操作划分的两个部分。

```golang
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
    
    
    // 递归操作
	left_slic = quickSort(left_slic)
	right_slic = quickSort(right_slic)


    // 合并数组
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



```

```
运行结果：

left: [1]
right: [2]
left: [1 2 3]
right: []
left: [1 2 3 4]
right: [5]
left: [7 8]
right: []
left: [7 8 9]
right: []
left: [7 8 9 10]
right: []
left: [1 2 3 4 5 6]
right: [7 8 9 10]
[1 2 3 4 5 6 7 8 9 10]

```

### 插入排序
在无序数组序列中，先选择第n个元素，并假设该元素前面（n-1）个元素的序列已排序，插入该序列使得排好顺序。

```golang
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


```

### 睡眠排序
通过给要排序的数组中的每个元素启动独立的任务，每个任务按照待排序元素的key执行相应的睡眠时间，然后按照时间将元素放到一起，达到排序的目的。
``` golang
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

```
