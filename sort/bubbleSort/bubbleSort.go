/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 6/22/20
 * Time: 22:18
 */
package main

import "fmt"

func main() {
	data := []int{3, 26, 0, 9, 6, 8, 123}
	sort := bubbleSort(data)
	fmt.Println("排序后", sort)
}

//冒泡排序
func bubbleSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}


func bubbleSort2(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := len(slice) - 1; j > i; j-- {
			if slice[j] < slice[j-1] {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}
		}
	}
	return slice
}
