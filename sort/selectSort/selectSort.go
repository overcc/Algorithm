/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 6/22/20
 * Time: 22:09
 */
package main

import "fmt"

func main() {
	data := []int{3, 26, 0, 9, 6, 8, 123}
	sort := selectSort(data)
	fmt.Println("排序后",sort)
}

//选择排序
func selectSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice)-1; j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

