package main

import (
	"fmt"
	"sort"
)

func arrayFunction(nums1 []int , nums2 []int) float64 {
	 nums3 := append(nums1 , nums2...)
	sort.Ints(nums3)
	 length := len(nums3)
	 if length %2 ==0 {
	 	result := float64(nums3[length/2-1]+(nums3[length/2]))/2
	 	fmt.Println(result)
	 }
	 return float64(nums3[length/2])

}


func main(){
	 nums1:= make([]int ,10)
	 nums2 := make([]int , 10)
	fmt.Println("enter first array numbers:")
	fmt.Scan(&nums1)
	fmt.Println("enter second  array numbers:")
	fmt.Scan(&nums2)
	 final:= arrayFunction(nums1, nums2)
fmt.Println(final)
}