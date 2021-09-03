package main

import "fmt"

 func lengthOfLongestSubstring(s string) int {
	  length , i  := 0,0
	  j:=0
	  value := make(map[byte]int)
	  str := value[s[i]]
	  str2 := value[s[j]]

	  for j = 0; j < len(s); j++ {
			str2 ++
		for str2 == 2 && i < j {
			str--
			i++
		}
		if length < j - i + 1 {
			length = j - i + 1
		}

      }
  return length
 }
func main(){
	var s string
	fmt.Println("print string:")
	fmt.Scan(&s)
	lengthOfLongestSubstring(s)
}