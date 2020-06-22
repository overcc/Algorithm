/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 6/22/20
 * Time: 22:45
 */
package main

import "fmt"

func main() {
	data := []int{3, 26, 0, 9, 6, 8, 123}
	sort := insertSort(data)
	fmt.Println("排序后", sort)
}

func insertSort(slice []int) []int {
	if slice == nil {
		return nil
	}
	for i := 1; i < len(slice); i++ {
		tmp, j := slice[i], i
		if slice[j-1] > tmp {
			for j >= 1 && slice[j-1] > tmp {
				slice[j] = slice[j-1]
				j--
			}

		}
		slice[j] = tmp
	}
	return slice
}
