package main
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
result := []int{}
l, r := 0, 0
for l < len(nums1) && r < len(nums2) {
if nums1[l] < nums2[r] {
result = append(result, nums1[l])
l++
} else {
result = append(result, nums2[r])
r++
}
}
result = append(result, nums1[l:]...)
result = append(result, nums2[r:]...)
if len(result)%2 == 0 {
return float64(result[len(result)/2-1]+result[len(result)/2]) / 2
}
return float64(result[len(result)/2])

